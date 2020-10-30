// DO NOT EDIT! Code generated.
package service

import (
	"context"
	"fmt"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/service"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/pipeproxy/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewServiceRefWithConfig)
	register.Register("def", NewServiceDefWithConfig)
	register.Register("none", newServiceNone)
}

type Config struct {
	Name string
	Def  service.Service `json:",omitempty"`
}

func NewServiceRefWithConfig(ctx context.Context, conf *Config) service.Service {
	o := &Service{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewServiceDefWithConfig(ctx context.Context, conf *Config) service.Service {
	return ServicePut(ctx, conf.Name, conf.Def)
}

func ServicePut(ctx context.Context, name string, def service.Service) service.Service {
	if def == nil {
		return ServiceNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return ServiceNone
	}
	store, _ := m.LoadOrStore("service.Service", map[string]service.Service{})
	store.(map[string]service.Service)[name] = def
	return def
}

func ServiceGet(ctx context.Context, name string, defaults service.Service) service.Service {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("service.Service")
		if ok {
			o, ok := store.(map[string]service.Service)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.Warnf("service.Service %q is not defined", name)
	return ServiceNone
}

var ServiceNone _ServiceNone

type _ServiceNone struct{}

func newServiceNone() service.Service {
	return ServiceNone
}

func (_ServiceNone) Close() (error error) {
	logger.Warn("this is none of service.Service")

	error = fmt.Errorf("error service.Service is none")

	return
}

func (_ServiceNone) Run(_ context.Context) (error error) {
	logger.Warn("this is none of service.Service")

	error = fmt.Errorf("error service.Service is none")

	return
}

type Service struct {
	Name string
	Def  service.Service
	Ctx  context.Context
}

func (o *Service) Close() error {
	return ServiceGet(o.Ctx, o.Name, o.Def).Close()
}

func (o *Service) Run(context context.Context) error {
	return ServiceGet(o.Ctx, o.Name, o.Def).Run(context)
}
