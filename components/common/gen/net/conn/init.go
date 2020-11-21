// DO NOT EDIT! Code generated.
package conn

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/wzshiming/logger"
)

func init() {
	register.Register("ref", NewConnRefWithConfig)
	register.Register("def", NewConnDefWithConfig)
	register.Register("none", newConnNone)
}

type Config struct {
	Name string
	Def  net.Conn `json:",omitempty"`
}

func NewConnRefWithConfig(ctx context.Context, conf *Config) net.Conn {
	o := &Conn{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewConnDefWithConfig(ctx context.Context, conf *Config) net.Conn {
	return ConnPut(ctx, conf.Name, conf.Def)
}

func ConnPut(ctx context.Context, name string, def net.Conn) net.Conn {
	if def == nil {
		return ConnNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return ConnNone
	}
	store, _ := m.LoadOrStore("net.Conn", map[string]net.Conn{})
	store.(map[string]net.Conn)[name] = def
	return def
}

func ConnGet(ctx context.Context, name string, defaults net.Conn) net.Conn {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, ok := m.Load("net.Conn")
		if ok {
			o, ok := store.(map[string]net.Conn)[name]
			if ok {
				return o
			}
		}
	}

	if defaults != nil {
		return defaults
	}
	logger.FromContext(ctx).V(-1).Info("net.Conn is not defined", "name", name)
	return ConnNone
}

var ConnNone _ConnNone

type _ConnNone struct{}

func newConnNone() net.Conn {
	return ConnNone
}

func (_ConnNone) Close() (error error) {
	logger.Log.V(-1).Info("this is none of net.Conn")

	error = fmt.Errorf("error net.Conn is none")

	return
}

func (_ConnNone) LocalAddr() (_ net.Addr) {
	logger.Log.V(-1).Info("this is none of net.Conn")

	return
}

func (_ConnNone) Read(_ []uint8) (_ int, error error) {
	logger.Log.V(-1).Info("this is none of net.Conn")

	error = fmt.Errorf("error net.Conn is none")

	return
}

func (_ConnNone) RemoteAddr() (_ net.Addr) {
	logger.Log.V(-1).Info("this is none of net.Conn")

	return
}

func (_ConnNone) SetDeadline(_ time.Time) (error error) {
	logger.Log.V(-1).Info("this is none of net.Conn")

	error = fmt.Errorf("error net.Conn is none")

	return
}

func (_ConnNone) SetReadDeadline(_ time.Time) (error error) {
	logger.Log.V(-1).Info("this is none of net.Conn")

	error = fmt.Errorf("error net.Conn is none")

	return
}

func (_ConnNone) SetWriteDeadline(_ time.Time) (error error) {
	logger.Log.V(-1).Info("this is none of net.Conn")

	error = fmt.Errorf("error net.Conn is none")

	return
}

func (_ConnNone) Write(_ []uint8) (_ int, error error) {
	logger.Log.V(-1).Info("this is none of net.Conn")

	error = fmt.Errorf("error net.Conn is none")

	return
}

type Conn struct {
	Name string
	Def  net.Conn
	Ctx  context.Context
}

func (o *Conn) Close() error {
	return ConnGet(o.Ctx, o.Name, o.Def).Close()
}

func (o *Conn) LocalAddr() net.Addr {
	return ConnGet(o.Ctx, o.Name, o.Def).LocalAddr()
}

func (o *Conn) Read(a []uint8) (int, error) {
	return ConnGet(o.Ctx, o.Name, o.Def).Read(a)
}

func (o *Conn) RemoteAddr() net.Addr {
	return ConnGet(o.Ctx, o.Name, o.Def).RemoteAddr()
}

func (o *Conn) SetDeadline(time time.Time) error {
	return ConnGet(o.Ctx, o.Name, o.Def).SetDeadline(time)
}

func (o *Conn) SetReadDeadline(time time.Time) error {
	return ConnGet(o.Ctx, o.Name, o.Def).SetReadDeadline(time)
}

func (o *Conn) SetWriteDeadline(time time.Time) error {
	return ConnGet(o.Ctx, o.Name, o.Def).SetWriteDeadline(time)
}

func (o *Conn) Write(a []uint8) (int, error) {
	return ConnGet(o.Ctx, o.Name, o.Def).Write(a)
}
