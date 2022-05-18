//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by ioc-go-cli

package rocketmq

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/normal"
)

func init() {
	normal.RegisterStructDescriber(&autowire.StructDescriber{
		Interface: new(RocketMQClient),
		Factory: func() interface{} {
			return &Impl{}
		},
		ParamFactory: func() interface{} {
			var _ ConfigInterface = &Config{}
			return &Config{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(ConfigInterface)
			impl := i.(*Impl)
			return param.New(impl)
		},
	})
}

type ConfigInterface interface {
	New(impl *Impl) (*Impl, error)
}
