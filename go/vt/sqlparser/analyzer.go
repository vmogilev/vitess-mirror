// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

// analyzer.go contains utility analysis functions.

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/youtube/vitess/go/sqltypes"
)

// These constants are used to identify the SQL statement type.
const (
	StmtSelect = iota
	StmtInsert
	StmtUpdate
	StmtDelete
	StmtDDL
	StmtBegin
	StmtCommit
	StmtRollback
	StmtSet
	StmtShow
	StmtUse
	StmtOther
	StmtUnknown
)

// Preview analyzes the beginning of the query and returns the
// statement type.
func Preview(sql string) int {
	trimmed := strings.TrimFunc(sql, unicode.IsSpace)
	prefix := trimmed
	if end := strings.IndexFunc(trimmed, unicode.IsSpace); end != -1 {
		prefix = trimmed[:end]
	}
	// Comparison is done in order of priority.
	lprefix := strings.ToLower(prefix)
	switch lprefix {
	case "select":
		return StmtSelect
	case "insert":
		return StmtInsert
	case "update":
		return StmtUpdate
	case "delete":
		return StmtDelete
	}
	switch strings.ToLower(trimmed) {
	case "begin", "start transaction":
		return StmtBegin
	case "commit":
		return StmtCommit
	case "rollback":
		return StmtRollback
	}
	switch lprefix {
	case "create", "alter", "rename", "drop":
		return StmtDDL
	case "set":
		return StmtSet
	case "show":
		return StmtShow
	case "use":
		return StmtUse
	case "analyze", "describe", "explain", "repair", "optimize", "truncate":
		return StmtOther
	}
	return StmtUnknown
}

// IsDML returns true if the query is an INSERT, UPDATE or DELETE statement.
func IsDML(sql string) bool {
	switch Preview(sql) {
	case StmtInsert, StmtUpdate, StmtDelete:
		return true
	}
	return false
}

// GetTableName returns the table name from the SimpleTableExpr
// only if it's a simple expression. Otherwise, it returns "".
func GetTableName(node SimpleTableExpr) TableIdent {
	if n, ok := node.(*TableName); ok && n.Qualifier.IsEmpty() {
		return n.Name
	}
	// sub-select or '.' expression
	return NewTableIdent("")
}

// IsColName returns true if the Expr is a *ColName.
func IsColName(node Expr) bool {
	_, ok := node.(*ColName)
	return ok
}

// IsValue returns true if the Expr is a string, integral or value arg.
// NULL is not considered to be a value.
func IsValue(node Expr) bool {
	switch v := node.(type) {
	case *SQLVal:
		switch v.Type {
		case StrVal, HexVal, IntVal, ValArg:
			return true
		}
	case *ValuesFuncExpr:
		if v.Resolved != nil {
			return IsValue(v.Resolved)
		}
	}
	return false
}

// IsNull returns true if the Expr is SQL NULL
func IsNull(node Expr) bool {
	switch node.(type) {
	case *NullVal:
		return true
	}
	return false
}

// IsSimpleTuple returns true if the Expr is a ValTuple that
// contains simple values or if it's a list arg.
func IsSimpleTuple(node Expr) bool {
	switch vals := node.(type) {
	case ValTuple:
		for _, n := range vals {
			if !IsValue(n) {
				return false
			}
		}
		return true
	case ListArg:
		return true
	}
	// It's a subquery
	return false
}

// AsInterface converts the Expr to an interface. It converts
// ValTuple to []interface{}, ValArg to string, StrVal to sqltypes.String,
// IntVal to sqltypes.Numeric, NullVal to nil.
// Otherwise, it returns an error.
func AsInterface(node Expr) (interface{}, error) {
	switch node := node.(type) {
	case *ValuesFuncExpr:
		if node.Resolved != nil {
			return AsInterface(node.Resolved)
		}
	case ValTuple:
		vals := make([]interface{}, 0, len(node))
		for _, val := range node {
			v, err := AsInterface(val)
			if err != nil {
				return nil, err
			}
			vals = append(vals, v)
		}
		return vals, nil
	case *SQLVal:
		switch node.Type {
		case ValArg:
			return string(node.Val), nil
		case StrVal:
			return sqltypes.MakeString(node.Val), nil
		case HexVal:
			v, err := node.HexDecode()
			if err != nil {
				return nil, err
			}
			return sqltypes.MakeString(v), nil
		case IntVal:
			n, err := sqltypes.BuildIntegral(string(node.Val))
			if err != nil {
				return nil, fmt.Errorf("type mismatch: %s", err)
			}
			return n, nil
		}
	case ListArg:
		return string(node), nil
	case *NullVal:
		return nil, nil
	}
	return nil, fmt.Errorf("expression is too complex '%v'", String(node))
}

// StringIn is a convenience function that returns
// true if str matches any of the values.
func StringIn(str string, values ...string) bool {
	for _, val := range values {
		if str == val {
			return true
		}
	}
	return false
}

// ExtractSetNums returns a map of key-num pairs
// if the query is a SET statement. Otherwise, it returns an
// error.
func ExtractSetNums(sql string) (map[string]int64, error) {
	stmt, err := Parse(sql)
	if err != nil {
		return nil, err
	}
	setStmt, ok := stmt.(*Set)
	if !ok {
		return nil, fmt.Errorf("ast did not yield *sqlparser.Set: %T", stmt)
	}
	result := make(map[string]int64)
	for _, expr := range setStmt.Exprs {
		if expr.Name.Qualifier != nil {
			return nil, fmt.Errorf("invalid syntax: %v", String(expr.Name))
		}
		key := expr.Name.Name.Lowered()

		sqlval, ok := expr.Expr.(*SQLVal)
		if !ok {
			return nil, fmt.Errorf("invalid syntax: %s", String(expr.Expr))
		}
		if sqlval.Type != IntVal {
			return nil, fmt.Errorf("invalid value type: %v", String(expr.Expr))
		}
		num, err := strconv.ParseInt(string(sqlval.Val), 0, 64)
		if err != nil {
			return nil, err
		}
		result[key] = num
	}
	return result, nil
}
