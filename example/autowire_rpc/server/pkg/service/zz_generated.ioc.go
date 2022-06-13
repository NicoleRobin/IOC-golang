//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package service

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"github.com/alibaba/ioc-golang/example/autowire_rpc/server/pkg/dto"
	rpc_service "github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_service"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceStruct_{}
		},
	})
	rpc_service.RegisterStructDescriptor(&autowire.StructDescriptor{
		Alias: "github.com/alibaba/ioc-golang/example/autowire_rpc/server/pkg/service/api.ServiceStructIOCRPCClient",
		Factory: func() interface{} {
			return &ServiceStruct{}
		},
	})
}

type serviceStruct_ struct {
	GetUser_ func(name string, age int) (*dto.User, error)
}

func (s *serviceStruct_) GetUser(name string, age int) (*dto.User, error) {
	return s.GetUser_(name, age)
}

type ServiceStructIOCInterface interface {
	GetUser(name string, age int) (*dto.User, error)
}

func GetServiceStruct() (*ServiceStruct, error) {
	i, err := rpc_service.GetImpl(util.GetSDIDByStructPtr(new(ServiceStruct)))
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceStruct)
	return impl, nil
}
func GetServiceStructIOCInterface() (ServiceStructIOCInterface, error) {
	return GetServiceStruct()
}
