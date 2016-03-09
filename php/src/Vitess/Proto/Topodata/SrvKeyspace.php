<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: topodata.proto

namespace Vitess\Proto\Topodata {

  class SrvKeyspace extends \DrSlump\Protobuf\Message {

    /**  @var \Vitess\Proto\Topodata\SrvKeyspace\KeyspacePartition[]  */
    public $partitions = array();
    
    /**  @var string */
    public $sharding_column_name = null;
    
    /**  @var int - \Vitess\Proto\Topodata\KeyspaceIdType */
    public $sharding_column_type = null;
    
    /**  @var \Vitess\Proto\Topodata\SrvKeyspace\ServedFrom[]  */
    public $served_from = array();
    
    /**  @var int */
    public $split_shard_count = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'topodata.SrvKeyspace');

      // REPEATED MESSAGE partitions = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "partitions";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $f->reference = '\Vitess\Proto\Topodata\SrvKeyspace\KeyspacePartition';
      $descriptor->addField($f);

      // OPTIONAL STRING sharding_column_name = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "sharding_column_name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL ENUM sharding_column_type = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "sharding_column_type";
      $f->type      = \DrSlump\Protobuf::TYPE_ENUM;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\Vitess\Proto\Topodata\KeyspaceIdType';
      $descriptor->addField($f);

      // REPEATED MESSAGE served_from = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "served_from";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $f->reference = '\Vitess\Proto\Topodata\SrvKeyspace\ServedFrom';
      $descriptor->addField($f);

      // OPTIONAL INT32 split_shard_count = 5
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 5;
      $f->name      = "split_shard_count";
      $f->type      = \DrSlump\Protobuf::TYPE_INT32;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <partitions> has a value
     *
     * @return boolean
     */
    public function hasPartitions(){
      return $this->_has(1);
    }
    
    /**
     * Clear <partitions> value
     *
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function clearPartitions(){
      return $this->_clear(1);
    }
    
    /**
     * Get <partitions> value
     *
     * @param int $idx
     * @return \Vitess\Proto\Topodata\SrvKeyspace\KeyspacePartition
     */
    public function getPartitions($idx = NULL){
      return $this->_get(1, $idx);
    }
    
    /**
     * Set <partitions> value
     *
     * @param \Vitess\Proto\Topodata\SrvKeyspace\KeyspacePartition $value
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function setPartitions(\Vitess\Proto\Topodata\SrvKeyspace\KeyspacePartition $value, $idx = NULL){
      return $this->_set(1, $value, $idx);
    }
    
    /**
     * Get all elements of <partitions>
     *
     * @return \Vitess\Proto\Topodata\SrvKeyspace\KeyspacePartition[]
     */
    public function getPartitionsList(){
     return $this->_get(1);
    }
    
    /**
     * Add a new element to <partitions>
     *
     * @param \Vitess\Proto\Topodata\SrvKeyspace\KeyspacePartition $value
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function addPartitions(\Vitess\Proto\Topodata\SrvKeyspace\KeyspacePartition $value){
     return $this->_add(1, $value);
    }
    
    /**
     * Check if <sharding_column_name> has a value
     *
     * @return boolean
     */
    public function hasShardingColumnName(){
      return $this->_has(2);
    }
    
    /**
     * Clear <sharding_column_name> value
     *
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function clearShardingColumnName(){
      return $this->_clear(2);
    }
    
    /**
     * Get <sharding_column_name> value
     *
     * @return string
     */
    public function getShardingColumnName(){
      return $this->_get(2);
    }
    
    /**
     * Set <sharding_column_name> value
     *
     * @param string $value
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function setShardingColumnName( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <sharding_column_type> has a value
     *
     * @return boolean
     */
    public function hasShardingColumnType(){
      return $this->_has(3);
    }
    
    /**
     * Clear <sharding_column_type> value
     *
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function clearShardingColumnType(){
      return $this->_clear(3);
    }
    
    /**
     * Get <sharding_column_type> value
     *
     * @return int - \Vitess\Proto\Topodata\KeyspaceIdType
     */
    public function getShardingColumnType(){
      return $this->_get(3);
    }
    
    /**
     * Set <sharding_column_type> value
     *
     * @param int - \Vitess\Proto\Topodata\KeyspaceIdType $value
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function setShardingColumnType( $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <served_from> has a value
     *
     * @return boolean
     */
    public function hasServedFrom(){
      return $this->_has(4);
    }
    
    /**
     * Clear <served_from> value
     *
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function clearServedFrom(){
      return $this->_clear(4);
    }
    
    /**
     * Get <served_from> value
     *
     * @param int $idx
     * @return \Vitess\Proto\Topodata\SrvKeyspace\ServedFrom
     */
    public function getServedFrom($idx = NULL){
      return $this->_get(4, $idx);
    }
    
    /**
     * Set <served_from> value
     *
     * @param \Vitess\Proto\Topodata\SrvKeyspace\ServedFrom $value
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function setServedFrom(\Vitess\Proto\Topodata\SrvKeyspace\ServedFrom $value, $idx = NULL){
      return $this->_set(4, $value, $idx);
    }
    
    /**
     * Get all elements of <served_from>
     *
     * @return \Vitess\Proto\Topodata\SrvKeyspace\ServedFrom[]
     */
    public function getServedFromList(){
     return $this->_get(4);
    }
    
    /**
     * Add a new element to <served_from>
     *
     * @param \Vitess\Proto\Topodata\SrvKeyspace\ServedFrom $value
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function addServedFrom(\Vitess\Proto\Topodata\SrvKeyspace\ServedFrom $value){
     return $this->_add(4, $value);
    }
    
    /**
     * Check if <split_shard_count> has a value
     *
     * @return boolean
     */
    public function hasSplitShardCount(){
      return $this->_has(5);
    }
    
    /**
     * Clear <split_shard_count> value
     *
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function clearSplitShardCount(){
      return $this->_clear(5);
    }
    
    /**
     * Get <split_shard_count> value
     *
     * @return int
     */
    public function getSplitShardCount(){
      return $this->_get(5);
    }
    
    /**
     * Set <split_shard_count> value
     *
     * @param int $value
     * @return \Vitess\Proto\Topodata\SrvKeyspace
     */
    public function setSplitShardCount( $value){
      return $this->_set(5, $value);
    }
  }
}

