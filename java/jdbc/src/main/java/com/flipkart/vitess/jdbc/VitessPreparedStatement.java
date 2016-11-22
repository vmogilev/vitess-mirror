package com.flipkart.vitess.jdbc;

import com.flipkart.vitess.util.Constants;
import com.flipkart.vitess.util.StringUtils;
import com.youtube.vitess.client.Context;
import com.youtube.vitess.client.VTGateConn;
import com.youtube.vitess.client.VTGateTx;
import com.youtube.vitess.client.cursor.Cursor;
import com.youtube.vitess.mysql.DateTime;
import com.youtube.vitess.proto.Topodata;

import java.io.InputStream;
import java.io.Reader;
import java.math.BigDecimal;
import java.net.URL;
import java.sql.Array;
import java.sql.BatchUpdateException;
import java.sql.Blob;
import java.sql.Clob;
import java.sql.Date;
import java.sql.NClob;
import java.sql.ParameterMetaData;
import java.sql.PreparedStatement;
import java.sql.Ref;
import java.sql.ResultSet;
import java.sql.ResultSetMetaData;
import java.sql.RowId;
import java.sql.SQLException;
import java.sql.SQLFeatureNotSupportedException;
import java.sql.SQLRecoverableException;
import java.sql.SQLXML;
import java.sql.Statement;
import java.sql.Time;
import java.sql.Timestamp;
import java.sql.Types;
import java.text.DateFormat;
import java.text.ParsePosition;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Calendar;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Locale;
import java.util.Map;
import java.util.logging.Logger;

/**
 * Created by harshit.gangal on 25/01/16.
 *
 * This class expected for an sql query and a given set of parameters
 * the DB Call can be made once with any of the following method
 * execute, executeQuery, executeUpdate and executeBatch.
 * After the call, the parameters will be reset
 * and a new set of parameters needs to be provided before calling any of the above method.
 */
public class VitessPreparedStatement extends VitessStatement implements PreparedStatement {

    /* Get actual class name to be printed on */
    private static Logger logger = Logger.getLogger(VitessPreparedStatement.class.getName());
    private final String sql;
    private final Map<String, Object> bindVariables;
    /** Holds batched commands */
    private List<Map<String, Object>> batchedArgs;

    public VitessPreparedStatement(VitessConnection vitessConnection, String sql)
        throws SQLException {
        this(vitessConnection, sql, ResultSet.TYPE_FORWARD_ONLY, ResultSet.CONCUR_READ_ONLY);
    }

    public VitessPreparedStatement(VitessConnection vitessConnection, String sql, int resultSetType,
        int resultSetConcurrency) throws SQLException {
        this(vitessConnection, sql, resultSetType, resultSetConcurrency,
            Statement.NO_GENERATED_KEYS);
    }

    public VitessPreparedStatement(VitessConnection vitessConnection, String sql,
        int autoGeneratedKeys) throws SQLException {
        this(vitessConnection, sql, ResultSet.TYPE_FORWARD_ONLY, ResultSet.CONCUR_READ_ONLY,
            autoGeneratedKeys);
    }

    public VitessPreparedStatement(VitessConnection vitessConnection, String sql, int resultSetType,
        int resultSetConcurrency, int autoGeneratedKeys) throws SQLException {
        super(vitessConnection, resultSetType, resultSetConcurrency);
        checkSQLNullOrEmpty(sql);
        this.bindVariables = new HashMap<>();
        this.sql = sql;
        this.generatedId = -1;
        this.retrieveGeneratedKeys = (autoGeneratedKeys == Statement.RETURN_GENERATED_KEYS);
        this.batchedArgs = new ArrayList<>();
    }

    public ResultSet executeQuery() throws SQLException {
        VTGateConn vtGateConn;
        Topodata.TabletType tabletType;
        Cursor cursor;
        boolean showSql;

        checkOpen();
        closeOpenResultSetAndResetCount();

        //Setting to default value
        this.generatedId = -1;

        vtGateConn = this.vitessConnection.getVtGateConn();
        tabletType = this.vitessConnection.getTabletType();

        showSql = StringUtils.startsWithIgnoreCaseAndWs(this.sql, Constants.SQL_SHOW);
        try {
            if (showSql) {
                cursor = this.executeShow(this.sql);
            } else {
                if (tabletType != Topodata.TabletType.MASTER || this.vitessConnection
                    .getAutoCommit()) {
                    Context context =
                        this.vitessConnection.createContext(this.queryTimeoutInMillis);
                    if (Constants.QueryExecuteType.SIMPLE == vitessConnection
                        .getExecuteTypeParam()) {
                        cursor =
                            vtGateConn.execute(context, this.sql, this.bindVariables, tabletType)
                                .checkedGet();
                    } else {
                        cursor = vtGateConn
                            .streamExecute(context, this.sql, this.bindVariables, tabletType);
                    }
                } else {
                    VTGateTx vtGateTx = this.vitessConnection.getVtGateTx();
                    if (vtGateTx == null) {
                        Context context =
                            this.vitessConnection.createContext(this.queryTimeoutInMillis);
                        vtGateTx = vtGateConn.begin(context).checkedGet();
                        this.vitessConnection.setVtGateTx(vtGateTx);
                    }
                    Context context =
                        this.vitessConnection.createContext(this.queryTimeoutInMillis);
                    cursor = vtGateTx.execute(context, this.sql, this.bindVariables, tabletType)
                        .checkedGet();
                }
            }

            if (null == cursor) {
                throw new SQLException(Constants.SQLExceptionMessages.METHOD_CALL_FAILED);
            }

            this.vitessResultSet = new VitessResultSet(cursor, this);
        } catch (SQLRecoverableException ex) {
            this.vitessConnection.setVtGateTx(null);
            throw ex;
        } finally {
            this.bindVariables.clear();
        }
        return (this.vitessResultSet);
    }

    public int executeUpdate() throws SQLException {
        VTGateConn vtGateConn;
        Topodata.TabletType tabletType;
        Cursor cursor;
        int truncatedUpdateCount;

        checkOpen();
        closeOpenResultSetAndResetCount();

        vtGateConn = this.vitessConnection.getVtGateConn();
        tabletType = this.vitessConnection.getTabletType();

        if (tabletType != Topodata.TabletType.MASTER) {
            throw new SQLException(Constants.SQLExceptionMessages.DML_NOT_ON_MASTER);
        }

        try {
            if (this.vitessConnection.getAutoCommit()) {
                Context context = this.vitessConnection.createContext(this.queryTimeoutInMillis);
                cursor = vtGateConn.execute(context, this.sql, this.bindVariables, tabletType)
                    .checkedGet();
            } else {
                VTGateTx vtGateTx = this.vitessConnection.getVtGateTx();
                if (null == vtGateTx) {
                    Context context =
                        this.vitessConnection.createContext(this.queryTimeoutInMillis);
                    vtGateTx = vtGateConn.begin(context).checkedGet();
                    this.vitessConnection.setVtGateTx(vtGateTx);
                }

                Context context = this.vitessConnection.createContext(this.queryTimeoutInMillis);
                cursor = vtGateTx.execute(context, this.sql, this.bindVariables, tabletType)
                    .checkedGet();
            }

            if (null == cursor) {
                throw new SQLException(Constants.SQLExceptionMessages.METHOD_CALL_FAILED);
            }

            if (!(null == cursor.getFields() || cursor.getFields().isEmpty())) {
                throw new SQLException(Constants.SQLExceptionMessages.SQL_RETURNED_RESULT_SET);
            }

            if (this.retrieveGeneratedKeys) {
                this.generatedId = cursor.getInsertId();
            }

            this.resultCount = cursor.getRowsAffected();

            if (this.resultCount > Integer.MAX_VALUE) {
                truncatedUpdateCount = Integer.MAX_VALUE;
            } else {
                truncatedUpdateCount = (int) this.resultCount;
            }
        } catch (SQLRecoverableException ex) {
            this.vitessConnection.setVtGateTx(null);
            throw ex;
        } finally {
            this.bindVariables.clear();
        }
        return truncatedUpdateCount;
    }

    public boolean execute() throws SQLException {
        Cursor cursor;
        boolean selectSql;
        boolean showSql;

        checkOpen();
        closeOpenResultSetAndResetCount();

        selectSql = StringUtils.startsWithIgnoreCaseAndWs(this.sql, Constants.SQL_SELECT);
        showSql = StringUtils.startsWithIgnoreCaseAndWs(this.sql, Constants.SQL_SHOW);

        if (showSql) {
            cursor = this.executeShow(this.sql);
            if (!(null == cursor || null == cursor.getFields() || cursor.getFields().isEmpty())) {
                this.vitessResultSet = new VitessResultSet(cursor, this);
                return true;
            }
            throw new SQLException(Constants.SQLExceptionMessages.METHOD_CALL_FAILED);
        } else if (selectSql) {
            this.executeQuery();
            return true;
        } else {
            this.executeUpdate();
            return false;
        }
    }

    public void clearParameters() throws SQLException {
        checkOpen();
        this.bindVariables.clear();
    }

    public void setNull(int parameterIndex, int sqlType) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, null);
    }

    public void setBoolean(int parameterIndex, boolean x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setByte(int parameterIndex, byte x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setShort(int parameterIndex, short x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setInt(int parameterIndex, int x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setLong(int parameterIndex, long x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setFloat(int parameterIndex, float x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setDouble(int parameterIndex, double x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setBigDecimal(int parameterIndex, BigDecimal x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setString(int parameterIndex, String x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setBytes(int parameterIndex, byte[] x) throws SQLException {
        checkOpen();
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, x);
    }

    public void setDate(int parameterIndex, Date x) throws SQLException {
        checkOpen();
        String date = DateTime.formatDate(x);
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, date);
    }

    public void setTime(int parameterIndex, Time x) throws SQLException {
        checkOpen();
        String time = DateTime.formatTime(x);
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, time);
    }

    public void setTimestamp(int parameterIndex, Timestamp x) throws SQLException {
        checkOpen();
        String timeStamp = DateTime.formatTimestamp(x);
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, timeStamp);
    }

    public void setDate(int parameterIndex, Date x, Calendar cal) throws SQLException {
        checkOpen();
        String date = DateTime.formatDate(x, cal);
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, date);
    }

    public void setTime(int parameterIndex, Time x, Calendar cal) throws SQLException {
        checkOpen();
        String time = DateTime.formatTime(x, cal);
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, time);
    }

    public void setTimestamp(int parameterIndex, Timestamp x, Calendar cal) throws SQLException {
        checkOpen();
        String timeStamp = DateTime.formatTimestamp(x, cal);
        this.bindVariables.put(Constants.LITERAL_V + parameterIndex, timeStamp);
    }

    public void setObject(int parameterIndex, Object x) throws SQLException {
        if (x == null) {
            setNull(parameterIndex, Types.NULL);
        } else if (x instanceof String) {
            setString(parameterIndex, (String) x);
        } else if (x instanceof Short) {
            setShort(parameterIndex, (Short) x);
        } else if (x instanceof Integer) {
            setInt(parameterIndex, (Integer) x);
        } else if (x instanceof Long) {
            setLong(parameterIndex, (Long) x);
        } else if (x instanceof Float) {
            setFloat(parameterIndex, (Float) x);
        } else if (x instanceof Double) {
            setDouble(parameterIndex, (Double) x);
        } else if (x instanceof Boolean) {
            setBoolean(parameterIndex, (Boolean) x);
        } else if (x instanceof Byte) {
            setByte(parameterIndex, (Byte) x);
        } else if (x instanceof Character) {
            setString(parameterIndex, String.valueOf(x));
        } else if (x instanceof Date) {
            setDate(parameterIndex, (Date) x);
        } else if (x instanceof Time) {
            setTime(parameterIndex, (Time) x);
        } else if (x instanceof Timestamp) {
            setTimestamp(parameterIndex, (Timestamp) x);
        } else if (x instanceof BigDecimal) {
            setBigDecimal(parameterIndex, (BigDecimal) x);
        } else {
            throw new SQLException(
                Constants.SQLExceptionMessages.SQL_TYPE_INFER + x.getClass().getCanonicalName());
        }
    }

    /**
     * Add bindVariables to the batch and clear it to have new set of bindVariables.
     *
     * @throws SQLException
     */
    public void addBatch() throws SQLException {
        checkOpen();
        this.batchedArgs.add(new HashMap<>(this.bindVariables));
        this.bindVariables.clear();
    }

    /**
     * Clear all the batched bindVariables.
     *
     * @throws SQLException
     */
    @Override public void clearBatch() throws SQLException {
        checkOpen();
        this.batchedArgs.clear();
    }

    /**
     * Execute the content in the batch.
     *
     * @return
     * @throws SQLException
     */
    @Override public int[] executeBatch() throws SQLException {
        checkOpen();
        int[] updateCounts = null;
        SQLException sqlEx = null;
        try {
            int numCommands = this.batchedArgs.size();
            if (numCommands > 0) {
                updateCounts = new int[numCommands];

                for (int i = 0; i < numCommands; i++) {
                    updateCounts[i] = Statement.EXECUTE_FAILED;
                }

                for (int commandIndex = 0; commandIndex < numCommands; commandIndex++) {
                    try {
                        Map<String, Object> bindVars = this.batchedArgs.get(commandIndex);
                        this.bindVariables.clear();
                        this.bindVariables.putAll(bindVars);
                        //TODO(harshit): To Support AutoGenerated Keys in a Batch
                        updateCounts[commandIndex] = this.executeUpdate();
                    } catch (SQLException ex) {
                        updateCounts[commandIndex] = Statement.EXECUTE_FAILED;
                        this.checkErrorAndReturn(ex, commandIndex);
                        sqlEx = ex;
                    }
                }
                if (sqlEx != null) {
                    throw new BatchUpdateException(sqlEx.getMessage(), sqlEx.getSQLState(),
                        sqlEx.getErrorCode(), updateCounts);
                }
            }
            return (updateCounts != null) ? updateCounts : new int[0];
        } finally {
            this.clearBatch();
        }
    }

    //Methods which are currently not supported

    public ParameterMetaData getParameterMetaData() throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setNull(int parameterIndex, int sqlType, String typeName) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setAsciiStream(int parameterIndex, InputStream x, int length) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setBinaryStream(int parameterIndex, InputStream x) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setCharacterStream(int parameterIndex, Reader reader, int length)
        throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setObject(int parameterIndex, Object parameterObject, int targetSqlType,
        int scaleOrLength) throws SQLException {
        if (null == parameterObject) {
            setNull(parameterIndex, Types.OTHER);
        } else {
            try {
                switch (targetSqlType) {
                    case Types.BOOLEAN:
                        if (parameterObject instanceof Boolean) {
                            setBoolean(parameterIndex, (Boolean) parameterObject);
                            break;
                        } else if (parameterObject instanceof String) {
                            setBoolean(parameterIndex,
                                "true".equalsIgnoreCase((String) parameterObject) || !"0"
                                    .equalsIgnoreCase((String) parameterObject));
                            break;
                        } else if (parameterObject instanceof Number) {
                            int intValue = ((Number) parameterObject).intValue();
                            setBoolean(parameterIndex, intValue != 0);
                            break;
                        } else {
                            throw new SQLException(
                                "Conversion from" + parameterObject.getClass().getName()
                                    + "to Types.Boolean is not Possible");
                        }
                    case Types.BIT:
                    case Types.TINYINT:
                    case Types.SMALLINT:
                    case Types.INTEGER:
                    case Types.BIGINT:
                    case Types.REAL:
                    case Types.FLOAT:
                    case Types.DOUBLE:
                    case Types.DECIMAL:
                    case Types.NUMERIC:
                        setNumericObject(parameterIndex, parameterObject, targetSqlType,
                            scaleOrLength);
                        break;
                    case Types.CHAR:
                    case Types.VARCHAR:
                    case Types.LONGVARCHAR:
                        if (parameterObject instanceof BigDecimal) {
                            setString(parameterIndex,
                                (StringUtils.fixDecimalExponent((parameterObject).toString())));
                        } else {
                            setString(parameterIndex, parameterObject.toString());
                        }
                        break;
                    case Types.CLOB:
                        if (parameterObject instanceof Clob) {
                            setClob(parameterIndex, (Clob) parameterObject);
                        } else {
                            setString(parameterIndex, parameterObject.toString());
                        }
                        break;
                    case Types.BINARY:
                    case Types.VARBINARY:
                    case Types.LONGVARBINARY:
                    case Types.BLOB:
                        if (parameterObject instanceof Blob) {
                            setBlob(parameterIndex, (Blob) parameterObject);
                        } else {
                            setBytes(parameterIndex, (byte[]) parameterObject);
                        }
                        break;
                    case Types.DATE:
                    case Types.TIMESTAMP:
                        java.util.Date parameterAsDate;
                        if (parameterObject instanceof String) {
                            ParsePosition pp = new ParsePosition(0);
                            DateFormat sdf = new SimpleDateFormat(
                                StringUtils.getDateTimePattern((String) parameterObject, false),
                                Locale.US);
                            parameterAsDate = sdf.parse((String) parameterObject, pp);
                        } else {
                            parameterAsDate = (java.util.Date) parameterObject;
                        }
                        switch (targetSqlType) {
                            case Types.DATE:
                                if (parameterAsDate instanceof Date) {
                                    setDate(parameterIndex, (Date) parameterAsDate);
                                } else {
                                    setDate(parameterIndex, new Date(parameterAsDate.getTime()));
                                }
                                break;
                            case Types.TIMESTAMP:
                                if (parameterAsDate instanceof Timestamp) {
                                    setTimestamp(parameterIndex, (Timestamp) parameterAsDate);
                                } else {
                                    setTimestamp(parameterIndex,
                                        new Timestamp(parameterAsDate.getTime()));
                                }
                                break;
                        }
                        break;
                    case Types.TIME:
                        if (parameterObject instanceof String) {
                            DateFormat sdf = new SimpleDateFormat(
                                StringUtils.getDateTimePattern((String) parameterObject, true),
                                Locale.US);
                            setTime(parameterIndex,
                                new Time(sdf.parse((String) parameterObject).getTime()));
                        } else if (parameterObject instanceof Timestamp) {
                            Timestamp timestamp = (Timestamp) parameterObject;
                            setTime(parameterIndex, new Time(timestamp.getTime()));
                        } else {
                            setTime(parameterIndex, (Time) parameterObject);
                        }
                        break;
                    default:
                        throw new SQLFeatureNotSupportedException(
                            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
                }
            } catch (Exception ex) {
                throw new SQLException(ex);
            }
        }
    }

    public void setAsciiStream(int parameterIndex, InputStream x, long length) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setBinaryStream(int parameterIndex, InputStream x, long length)
        throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setCharacterStream(int parameterIndex, Reader reader, long length)
        throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setUnicodeStream(int parameterIndex, InputStream x, int length)
        throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setRef(int parameterIndex, Ref x) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setBlob(int parameterIndex, Blob x) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setClob(int parameterIndex, Clob x) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setArray(int parameterIndex, Array x) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public ResultSetMetaData getMetaData() throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setURL(int parameterIndex, URL x) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setRowId(int parameterIndex, RowId x) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setNString(int parameterIndex, String value) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setNCharacterStream(int parameterIndex, Reader value, long length)
        throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setNClob(int parameterIndex, NClob value) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setClob(int parameterIndex, Reader reader, long length) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setBlob(int parameterIndex, InputStream inputStream, long length)
        throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setNClob(int parameterIndex, Reader reader, long length) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setSQLXML(int parameterIndex, SQLXML xmlObject) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setAsciiStream(int parameterIndex, InputStream x) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setBinaryStream(int parameterIndex, InputStream x, int length) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setCharacterStream(int parameterIndex, Reader reader) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setNCharacterStream(int parameterIndex, Reader value) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setClob(int parameterIndex, Reader reader) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setBlob(int parameterIndex, InputStream inputStream) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setNClob(int parameterIndex, Reader reader) throws SQLException {
        throw new SQLFeatureNotSupportedException(
            Constants.SQLExceptionMessages.SQL_FEATURE_NOT_SUPPORTED);
    }

    public void setObject(int parameterIndex, Object parameterObject, int targetSqlType)
        throws SQLException {
        if (!(parameterObject instanceof BigDecimal)) {
            setObject(parameterIndex, parameterObject, targetSqlType, 0);
        } else {
            setObject(parameterIndex, parameterObject, targetSqlType,
                ((BigDecimal) parameterObject).scale());
        }
    }

    private void setNumericObject(int parameterIndex, Object parameterObj, int targetSqlType,
        int scale) throws SQLException {
        Number numberParam;
        if (parameterObj instanceof Boolean) {
            numberParam = (Boolean) parameterObj ? Integer.valueOf(1) : Integer.valueOf(0);
        } else if (parameterObj instanceof String) {
            switch (targetSqlType) {
                case Types.BIT:
                    if ("1".equals(parameterObj) || "0".equals(parameterObj)) {
                        numberParam = Integer.valueOf((String) parameterObj);
                    } else {
                        boolean parameterAsBoolean = "true".equalsIgnoreCase((String) parameterObj);
                        numberParam = parameterAsBoolean ? Integer.valueOf(1) : Integer.valueOf(0);
                    }
                    break;

                case Types.TINYINT:
                case Types.SMALLINT:
                case Types.INTEGER:
                    numberParam = Integer.valueOf((String) parameterObj);
                    break;

                case Types.BIGINT:
                    numberParam = Long.valueOf((String) parameterObj);
                    break;

                case Types.REAL:
                    numberParam = Float.valueOf((String) parameterObj);
                    break;

                case Types.FLOAT:
                case Types.DOUBLE:
                    numberParam = Double.valueOf((String) parameterObj);
                    break;

                case Types.DECIMAL:
                case Types.NUMERIC:
                default:
                    numberParam = new java.math.BigDecimal((String) parameterObj);
            }
        } else {
            numberParam = (Number) parameterObj;
        }
        switch (targetSqlType) {
            case Types.BIT:
            case Types.TINYINT:
            case Types.SMALLINT:
            case Types.INTEGER:
                setInt(parameterIndex, numberParam.intValue());
                break;

            case Types.BIGINT:
                setLong(parameterIndex, numberParam.longValue());
                break;

            case Types.REAL:
                setFloat(parameterIndex, numberParam.floatValue());
                break;

            case Types.FLOAT:
            case Types.DOUBLE:
                setDouble(parameterIndex, numberParam.doubleValue());
                break;

            case Types.DECIMAL:
            case Types.NUMERIC:

                if (numberParam instanceof java.math.BigDecimal) {
                    BigDecimal scaledBigDecimal;
                    try {
                        scaledBigDecimal = ((java.math.BigDecimal) numberParam).setScale(scale);
                    } catch (ArithmeticException ex) {
                        try {
                            scaledBigDecimal = ((java.math.BigDecimal) numberParam)
                                .setScale(scale, BigDecimal.ROUND_HALF_UP);
                        } catch (ArithmeticException arEx) {
                            throw new SQLException(
                                "Can't set the scale of '" + scale + "' for Decimal Argument"
                                    + numberParam);
                        }
                    }
                    setBigDecimal(parameterIndex, scaledBigDecimal);
                } else if (numberParam instanceof java.math.BigInteger) {
                    setBigDecimal(parameterIndex,
                        new java.math.BigDecimal((java.math.BigInteger) numberParam, scale));
                } else {
                    setBigDecimal(parameterIndex,
                        new java.math.BigDecimal(numberParam.doubleValue()));
                }
                break;
        }
    }

    /**
     * Get the batched args as added by the addBatch method(s).
     * The list is unmodifiable and will contain BatchedBindValues.
     *
     * @return an unmodifiable List of batched args
     */
    @Override public List getBatchedArgs() {
        return Collections.unmodifiableList(this.batchedArgs);
    }
}
