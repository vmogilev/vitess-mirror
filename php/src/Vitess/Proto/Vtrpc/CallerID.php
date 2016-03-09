<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: vtrpc.proto

namespace Vitess\Proto\Vtrpc {

  class CallerID extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $principal = null;
    
    /**  @var string */
    public $component = null;
    
    /**  @var string */
    public $subcomponent = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'vtrpc.CallerID');

      // OPTIONAL STRING principal = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "principal";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING component = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "component";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING subcomponent = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "subcomponent";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <principal> has a value
     *
     * @return boolean
     */
    public function hasPrincipal(){
      return $this->_has(1);
    }
    
    /**
     * Clear <principal> value
     *
     * @return \Vitess\Proto\Vtrpc\CallerID
     */
    public function clearPrincipal(){
      return $this->_clear(1);
    }
    
    /**
     * Get <principal> value
     *
     * @return string
     */
    public function getPrincipal(){
      return $this->_get(1);
    }
    
    /**
     * Set <principal> value
     *
     * @param string $value
     * @return \Vitess\Proto\Vtrpc\CallerID
     */
    public function setPrincipal( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <component> has a value
     *
     * @return boolean
     */
    public function hasComponent(){
      return $this->_has(2);
    }
    
    /**
     * Clear <component> value
     *
     * @return \Vitess\Proto\Vtrpc\CallerID
     */
    public function clearComponent(){
      return $this->_clear(2);
    }
    
    /**
     * Get <component> value
     *
     * @return string
     */
    public function getComponent(){
      return $this->_get(2);
    }
    
    /**
     * Set <component> value
     *
     * @param string $value
     * @return \Vitess\Proto\Vtrpc\CallerID
     */
    public function setComponent( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <subcomponent> has a value
     *
     * @return boolean
     */
    public function hasSubcomponent(){
      return $this->_has(3);
    }
    
    /**
     * Clear <subcomponent> value
     *
     * @return \Vitess\Proto\Vtrpc\CallerID
     */
    public function clearSubcomponent(){
      return $this->_clear(3);
    }
    
    /**
     * Get <subcomponent> value
     *
     * @return string
     */
    public function getSubcomponent(){
      return $this->_get(3);
    }
    
    /**
     * Set <subcomponent> value
     *
     * @param string $value
     * @return \Vitess\Proto\Vtrpc\CallerID
     */
    public function setSubcomponent( $value){
      return $this->_set(3, $value);
    }
  }
}

