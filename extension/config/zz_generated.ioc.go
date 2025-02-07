//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package config

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	autowireconfig "github.com/alibaba/ioc-golang/extension/autowire/config"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configFloat64_{}
		},
	})
	configFloat64StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigFloat64)
		},
		ParamFactory: func() interface{} {
			return new(ConfigFloat64)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configFloat64Interface)
			impl := i.(*ConfigFloat64)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configFloat64StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configInt64_{}
		},
	})
	configInt64StructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigInt64)
		},
		ParamFactory: func() interface{} {
			return new(ConfigInt64)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configInt64Interface)
			impl := i.(*ConfigInt64)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configInt64StructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configInt_{}
		},
	})
	configIntStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigInt)
		},
		ParamFactory: func() interface{} {
			return new(ConfigInt)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configIntInterface)
			impl := i.(*ConfigInt)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configIntStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configMap_{}
		},
	})
	configMapStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigMap)
		},
		ParamFactory: func() interface{} {
			return new(ConfigMap)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configMapInterface)
			impl := i.(*ConfigMap)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configMapStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configSlice_{}
		},
	})
	configSliceStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigSlice)
		},
		ParamFactory: func() interface{} {
			return new(ConfigSlice)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configSliceInterface)
			impl := i.(*ConfigSlice)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configSliceStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &configString_{}
		},
	})
	configStringStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return new(ConfigString)
		},
		ParamFactory: func() interface{} {
			return new(ConfigString)
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configStringInterface)
			impl := i.(*ConfigString)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	autowireconfig.RegisterStructDescriptor(configStringStructDescriptor)
}

type configStringInterface interface {
	new(impl *ConfigString) (*ConfigString, error)
}
type configFloat64Interface interface {
	new(impl *ConfigFloat64) (*ConfigFloat64, error)
}
type configInt64Interface interface {
	new(impl *ConfigInt64) (*ConfigInt64, error)
}
type configIntInterface interface {
	new(impl *ConfigInt) (*ConfigInt, error)
}
type configMapInterface interface {
	new(impl *ConfigMap) (*ConfigMap, error)
}
type configSliceInterface interface {
	new(impl *ConfigSlice) (*ConfigSlice, error)
}
type configFloat64_ struct {
	Value_ func() float64
	new_   func(impl *ConfigFloat64) (*ConfigFloat64, error)
}

func (c *configFloat64_) Value() float64 {
	return c.Value_()
}

func (c *configFloat64_) new(impl *ConfigFloat64) (*ConfigFloat64, error) {
	return c.new_(impl)
}

type configInt64_ struct {
	Value_ func() int64
	new_   func(impl *ConfigInt64) (*ConfigInt64, error)
}

func (c *configInt64_) Value() int64 {
	return c.Value_()
}

func (c *configInt64_) new(impl *ConfigInt64) (*ConfigInt64, error) {
	return c.new_(impl)
}

type configInt_ struct {
	Value_ func() int
	new_   func(impl *ConfigInt) (*ConfigInt, error)
}

func (c *configInt_) Value() int {
	return c.Value_()
}

func (c *configInt_) new(impl *ConfigInt) (*ConfigInt, error) {
	return c.new_(impl)
}

type configMap_ struct {
	Value_ func() map[string]interface{}
	new_   func(impl *ConfigMap) (*ConfigMap, error)
}

func (c *configMap_) Value() map[string]interface{} {
	return c.Value_()
}

func (c *configMap_) new(impl *ConfigMap) (*ConfigMap, error) {
	return c.new_(impl)
}

type configSlice_ struct {
	Value_ func() []interface{}
	new_   func(impl *ConfigSlice) (*ConfigSlice, error)
}

func (c *configSlice_) Value() []interface{} {
	return c.Value_()
}

func (c *configSlice_) new(impl *ConfigSlice) (*ConfigSlice, error) {
	return c.new_(impl)
}

type configString_ struct {
	Value_ func() string
	new_   func(impl *ConfigString) (*ConfigString, error)
}

func (c *configString_) Value() string {
	return c.Value_()
}

func (c *configString_) new(impl *ConfigString) (*ConfigString, error) {
	return c.new_(impl)
}

type ConfigFloat64IOCInterface interface {
	Value() float64
	new(impl *ConfigFloat64) (*ConfigFloat64, error)
}

type ConfigInt64IOCInterface interface {
	Value() int64
	new(impl *ConfigInt64) (*ConfigInt64, error)
}

type ConfigIntIOCInterface interface {
	Value() int
	new(impl *ConfigInt) (*ConfigInt, error)
}

type ConfigMapIOCInterface interface {
	Value() map[string]interface{}
	new(impl *ConfigMap) (*ConfigMap, error)
}

type ConfigSliceIOCInterface interface {
	Value() []interface{}
	new(impl *ConfigSlice) (*ConfigSlice, error)
}

type ConfigStringIOCInterface interface {
	Value() string
	new(impl *ConfigString) (*ConfigString, error)
}

var _configFloat64SDID string
var _configInt64SDID string
var _configIntSDID string
var _configMapSDID string
var _configSliceSDID string
var _configStringSDID string
