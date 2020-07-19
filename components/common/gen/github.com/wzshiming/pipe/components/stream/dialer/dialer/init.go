// DO NOT EDIT! Code generated.
package reference

import (
	"context"
	"net"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stream/dialer"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewDialerRefWithConfig)
	register.Register("def", NewDialerDefWithConfig)
	register.Register("none", NewDialerNone)
}

type Config struct {
	Name string
	Def  dialer.Dialer `json:",omitempty"`
}

func NewDialerRefWithConfig(conf *Config) (dialer.Dialer, error) {
	o := &Dialer{
		Name: conf.Name,
		Def:  conf.Def,
	}
	return o, nil
}

func NewDialerDefWithConfig(conf *Config) (dialer.Dialer, error) {
	DialerStore[conf.Name] = conf.Def
	return conf.Def, nil
}

var DialerStore = map[string]dialer.Dialer{}

func DialerFind(name string, defaults dialer.Dialer) dialer.Dialer {
	o, ok := DialerStore[name]
	if ok {
		return o
	}
	if defaults != nil {
		return defaults
	}
	return DialerNone{}
}

type DialerNone struct{}

func NewDialerNone() dialer.Dialer {
	return DialerNone{}
}

func (DialerNone) DialStream(_ context.Context) (_ net.Conn, _ error) {
	logger.Warn("this is none of dialer.Dialer")
	return
}

type Dialer struct {
	Name string
	Def  dialer.Dialer
}

func (o *Dialer) DialStream(context context.Context) (net.Conn, error) {
	return DialerFind(o.Name, o.Def).DialStream(context)
}
