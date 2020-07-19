// DO NOT EDIT! Code generated.
package reference

import (
	"context"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/service"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewServiceRefWithConfig)
	register.Register("def", NewServiceDefWithConfig)
	register.Register("none", NewServiceNone)
}

type Config struct {
	Name string
	Def  service.Service `json:",omitempty"`
}

func NewServiceRefWithConfig(conf *Config) (service.Service, error) {
	o := &Service{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewServiceDefWithConfig(conf *Config) (service.Service, error) {
	ServiceStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var ServiceStore = map[string]service.Service{}

func ServiceFind(name string, defaults service.Service) service.Service {
	o, ok := ServiceStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return ServiceNone{}
}

type ServiceNone struct{}

func NewServiceNone() service.Service {
	return ServiceNone{}
}

func (ServiceNone) Close() (_ error) {
	logger.Warn("this is none of service.Service")
	return
}

func (ServiceNone) Run(_ context.Context) (_ error) {
	logger.Warn("this is none of service.Service")
	return
}

type Service struct {
	Name string
	Def  service.Service
}

func (o *Service) Close() error {
	return ServiceFind(o.Name, o.Def).Close()
}

func (o *Service) Run(context context.Context) error {
	return ServiceFind(o.Name, o.Def).Run(context)
}
