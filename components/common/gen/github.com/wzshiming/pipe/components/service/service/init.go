// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"fmt"
	"sync"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/service"
	"github.com/wzshiming/pipe/internal/logger"
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

func NewServiceRefWithConfig(conf *Config) service.Service {
	o := &Service{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o
}

func NewServiceDefWithConfig(conf *Config) service.Service {
	return ServicePut(conf.Name, conf.Def)
}

var (
	mut           sync.RWMutex
	_ServiceStore = map[string]service.Service{}
)

func ServicePut(name string, def service.Service) service.Service {
	if def == nil {
		def = ServiceNone
	}
	mut.Lock()
	_ServiceStore[name] = def
	mut.Unlock()
	return def
}

func ServiceGet(name string, defaults service.Service) service.Service {
	mut.RLock()
	o, ok := _ServiceStore[name]
	mut.RUnlock()
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return ServiceNone
}

var ServiceNone _ServiceNone

type _ServiceNone struct{}

func newServiceNone() service.Service {
	return ServiceNone
}

func (_ServiceNone) Close() (error error) {
	logger.Warn("this is none of service.Service")

	error = fmt.Errorf("error none")

	return
}

func (_ServiceNone) Run(_ context.Context) (error error) {
	logger.Warn("this is none of service.Service")

	error = fmt.Errorf("error none")

	return
}

type Service struct {
	Name string
	Def  service.Service
}

func (o *Service) Close() error {
	return ServiceGet(o.Name, o.Def).Close()
}

func (o *Service) Run(context context.Context) error {
	return ServiceGet(o.Name, o.Def).Run(context)
}
