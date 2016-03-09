<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: binlogservice.proto

namespace Vitess\Proto\Binlogservice {

  class UpdateStreamClient extends \Grpc\BaseStub {

    public function __construct($hostname, $opts) {
      parent::__construct($hostname, $opts);
    }
    /**
     * @param Vitess\Proto\Binlogdata\StreamUpdateRequest $input
     */
    public function StreamUpdate($argument, $metadata = array(), $options = array()) {
      return $this->_serverStreamRequest('/binlogservice.UpdateStream/StreamUpdate', $argument, '\Vitess\Proto\Binlogdata\StreamUpdateResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Binlogdata\StreamKeyRangeRequest $input
     */
    public function StreamKeyRange($argument, $metadata = array(), $options = array()) {
      return $this->_serverStreamRequest('/binlogservice.UpdateStream/StreamKeyRange', $argument, '\Vitess\Proto\Binlogdata\StreamKeyRangeResponse::deserialize', $metadata, $options);
    }
    /**
     * @param Vitess\Proto\Binlogdata\StreamTablesRequest $input
     */
    public function StreamTables($argument, $metadata = array(), $options = array()) {
      return $this->_serverStreamRequest('/binlogservice.UpdateStream/StreamTables', $argument, '\Vitess\Proto\Binlogdata\StreamTablesResponse::deserialize', $metadata, $options);
    }
  }
}
