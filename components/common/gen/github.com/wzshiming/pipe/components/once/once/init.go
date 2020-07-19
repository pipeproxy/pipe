// DO NOT EDIT! Code generated.
package reference

import (
	"context"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/once"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewOnceRefWithConfig)
	register.Register("def", NewOnceDefWithConfig)
	register.Register("none", NewOnceNone)
}

type Config struct {
	Name string
	Def  once.Once `json:",omitempty"`
}

func NewOnceRefWithConfig(conf *Config) (once.Once, error) {
	o := &Once{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewOnceDefWithConfig(conf *Config) (once.Once, error) {
	OnceStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var OnceStore = map[string]once.Once{}

func OnceFind(name string, defaults once.Once) once.Once {
	o, ok := OnceStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return OnceNone{}
}

type OnceNone struct{}

func NewOnceNone() once.Once {
	return OnceNone{}
}

func (OnceNone) Do(_ context.Context) (_ error) {
	logger.Warn("this is none of once.Once")
	return
}

type Once struct {
	Name string
	Def  once.Once
}

func (o *Once) Do(context context.Context) error {
	return OnceFind(o.Name, o.Def).Do(context)
}
