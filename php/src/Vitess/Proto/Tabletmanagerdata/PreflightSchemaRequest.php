<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: tabletmanagerdata.proto

namespace Vitess\Proto\Tabletmanagerdata {

  class PreflightSchemaRequest extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $change = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'tabletmanagerdata.PreflightSchemaRequest');

      // OPTIONAL STRING change = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "change";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <change> has a value
     *
     * @return boolean
     */
    public function hasChange(){
      return $this->_has(1);
    }
    
    /**
     * Clear <change> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\PreflightSchemaRequest
     */
    public function clearChange(){
      return $this->_clear(1);
    }
    
    /**
     * Get <change> value
     *
     * @return string
     */
    public function getChange(){
      return $this->_get(1);
    }
    
    /**
     * Set <change> value
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\PreflightSchemaRequest
     */
    public function setChange( $value){
      return $this->_set(1, $value);
    }
  }
}

