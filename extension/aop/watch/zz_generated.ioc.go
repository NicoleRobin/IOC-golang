//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package watch

import (
	"github.com/alibaba/ioc-golang/aop"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &interceptorImpl_{}
		},
	})
	interceptorImplStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &interceptorImpl{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	singleton.RegisterStructDescriptor(interceptorImplStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &context_{}
		},
	})
	contextStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &context{}
		},
		ParamFactory: func() interface{} {
			var _ contextParamInterface = &contextParam{}
			return &contextParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(contextParamInterface)
			impl := i.(*context)
			return param.new(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	normal.RegisterStructDescriptor(contextStructDescriptor)
	watchServiceStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &watchService{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	singleton.RegisterStructDescriptor(watchServiceStructDescriptor)
}

type contextParamInterface interface {
	new(impl *context) (*context, error)
}
type interceptorImpl_ struct {
	BeforeInvoke_ func(ctx *aop.InvocationContext)
	AfterInvoke_  func(ctx *aop.InvocationContext)
	Watch_        func(watchCtx contextIOCInterface)
	UnWatch_      func(watchCtx contextIOCInterface)
}

func (i *interceptorImpl_) BeforeInvoke(ctx *aop.InvocationContext) {
	i.BeforeInvoke_(ctx)
}

func (i *interceptorImpl_) AfterInvoke(ctx *aop.InvocationContext) {
	i.AfterInvoke_(ctx)
}

func (i *interceptorImpl_) Watch(watchCtx contextIOCInterface) {
	i.Watch_(watchCtx)
}

func (i *interceptorImpl_) UnWatch(watchCtx contextIOCInterface) {
	i.UnWatch_(watchCtx)
}

type context_ struct {
	getSDID_      func() string
	getMethod_    func() string
	beforeInvoke_ func(ctx *aop.InvocationContext)
	afterInvoke_  func(ctx *aop.InvocationContext)
}

func (c *context_) getSDID() string {
	return c.getSDID_()
}

func (c *context_) getMethod() string {
	return c.getMethod_()
}

func (c *context_) beforeInvoke(ctx *aop.InvocationContext) {
	c.beforeInvoke_(ctx)
}

func (c *context_) afterInvoke(ctx *aop.InvocationContext) {
	c.afterInvoke_(ctx)
}

type interceptorImplIOCInterface interface {
	BeforeInvoke(ctx *aop.InvocationContext)
	AfterInvoke(ctx *aop.InvocationContext)
	Watch(watchCtx contextIOCInterface)
	UnWatch(watchCtx contextIOCInterface)
}

type contextIOCInterface interface {
	getSDID() string
	getMethod() string
	beforeInvoke(ctx *aop.InvocationContext)
	afterInvoke(ctx *aop.InvocationContext)
}

var _interceptorImplSDID string

func GetinterceptorImplSingleton() (*interceptorImpl, error) {
	if _interceptorImplSDID == "" {
		_interceptorImplSDID = util.GetSDIDByStructPtr(new(interceptorImpl))
	}
	i, err := singleton.GetImpl(_interceptorImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*interceptorImpl)
	return impl, nil
}

func GetinterceptorImplIOCInterfaceSingleton() (interceptorImplIOCInterface, error) {
	if _interceptorImplSDID == "" {
		_interceptorImplSDID = util.GetSDIDByStructPtr(new(interceptorImpl))
	}
	i, err := singleton.GetImplWithProxy(_interceptorImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(interceptorImplIOCInterface)
	return impl, nil
}

type ThisinterceptorImpl struct {
}

func (t *ThisinterceptorImpl) This() interceptorImplIOCInterface {
	thisPtr, _ := GetinterceptorImplIOCInterfaceSingleton()
	return thisPtr
}

var _contextSDID string

func Getcontext(p *contextParam) (*context, error) {
	if _contextSDID == "" {
		_contextSDID = util.GetSDIDByStructPtr(new(context))
	}
	i, err := normal.GetImpl(_contextSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*context)
	return impl, nil
}

func GetcontextIOCInterface(p *contextParam) (contextIOCInterface, error) {
	if _contextSDID == "" {
		_contextSDID = util.GetSDIDByStructPtr(new(context))
	}
	i, err := normal.GetImplWithProxy(_contextSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(contextIOCInterface)
	return impl, nil
}

var _watchServiceSDID string

func GetwatchServiceSingleton() (*watchService, error) {
	if _watchServiceSDID == "" {
		_watchServiceSDID = util.GetSDIDByStructPtr(new(watchService))
	}
	i, err := singleton.GetImpl(_watchServiceSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*watchService)
	return impl, nil
}
