// DO NOT EDIT! Code generated.
package policy

import (
	"context"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/wzshiming/logger"
)

func init() {
	register.Register("ref", NewPolicyRefWithConfig)
	register.Register("def", NewPolicyDefWithConfig)
	register.Register("none", newPolicyNone)
}

type Config struct {
	Name string
	Def  balance.Policy `json:",omitempty"`
}

func NewPolicyRefWithConfig(ctx context.Context, conf *Config) balance.Policy {
	o := &Policy{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewPolicyDefWithConfig(ctx context.Context, conf *Config) balance.Policy {
	return PolicyPut(ctx, conf.Name, conf.Def)
}

func PolicyPut(ctx context.Context, name string, def balance.Policy) balance.Policy {
	if def == nil {
		return PolicyNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return PolicyNone
	}
	store, _ := m.LoadOrStore("balance.Policy", map[string]balance.Policy{})
	store.(map[string]balance.Policy)[name] = def
	return def
}

func PolicyGet(ctx context.Context, name string, defaults balance.Policy) balance.Policy {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("balance.Policy")
		if ok {
			o, ok := store.(map[string]balance.Policy)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.FromContext(ctx).V(-1).Info("balance.Policy is not defined", "name", name)
	return PolicyNone
}

var PolicyNone _PolicyNone

type _PolicyNone struct{}

func newPolicyNone() balance.Policy {
	return PolicyNone
}

func (_PolicyNone) Clone() (_ balance.Policy) {
	logger.Log.V(-1).Info("this is none of balance.Policy")

	return
}

func (_PolicyNone) InUse(_ func(uint64)) {
	logger.Log.V(-1).Info("this is none of balance.Policy")

	return
}

func (_PolicyNone) Init(_ uint64) {
	logger.Log.V(-1).Info("this is none of balance.Policy")

	return
}

type Policy struct {
	Name string
	Def  balance.Policy
	Ctx  context.Context
}

func (o *Policy) Clone() balance.Policy {
	return PolicyGet(o.Ctx, o.Name, o.Def).Clone()
}

func (o *Policy) InUse(a func(uint64)) {
	PolicyGet(o.Ctx, o.Name, o.Def).InUse(a)
}

func (o *Policy) Init(uint64 uint64) {
	PolicyGet(o.Ctx, o.Name, o.Def).Init(uint64)
}
