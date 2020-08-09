// DO NOT EDIT! Code generated.
package packetconn

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/internal/ctxcache"
	"github.com/wzshiming/pipe/internal/logger"
)

func init() {
	register.Register("ref", NewPacketConnRefWithConfig)
	register.Register("def", NewPacketConnDefWithConfig)
	register.Register("none", newPacketConnNone)
}

type Config struct {
	Name string
	Def  net.PacketConn `json:",omitempty"`
}

func NewPacketConnRefWithConfig(ctx context.Context, conf *Config) net.PacketConn {
	o := &PacketConn{
		Name: conf.Name,
		Def:  conf.Def,
		Ctx:  ctx,
	}
	return o
}

func NewPacketConnDefWithConfig(ctx context.Context, conf *Config) net.PacketConn {
	return PacketConnPut(ctx, conf.Name, conf.Def)
}

func PacketConnPut(ctx context.Context, name string, def net.PacketConn) net.PacketConn {
	if def == nil {
		def = PacketConnNone
	}

	m, ok := ctxcache.GetCacheWithContext(ctx)
	if !ok {
		return PacketConnNone
	}
	store, _ := m.LoadOrStore("net.PacketConn", map[string]net.PacketConn{})
	store.(map[string]net.PacketConn)[name] = def
	return def
}

func PacketConnGet(ctx context.Context, name string, defaults net.PacketConn) net.PacketConn {
	m, ok := ctxcache.GetCacheWithContext(ctx)
	if ok {
		store, _ := m.LoadOrStore("net.PacketConn", map[string]net.PacketConn{})
		o, ok := store.(map[string]net.PacketConn)[name]
		if ok {
			return o
		}
	}

	if defaults != nil {
		return defaults
	}
	return PacketConnNone
}

var PacketConnNone _PacketConnNone

type _PacketConnNone struct{}

func newPacketConnNone() net.PacketConn {
	return PacketConnNone
}

func (_PacketConnNone) Close() (error error) {
	logger.Warn("this is none of net.PacketConn")

	error = fmt.Errorf("error none")

	return
}

func (_PacketConnNone) LocalAddr() (_ net.Addr) {
	logger.Warn("this is none of net.PacketConn")

	return
}

func (_PacketConnNone) ReadFrom(_ []uint8) (_ int, _ net.Addr, error error) {
	logger.Warn("this is none of net.PacketConn")

	error = fmt.Errorf("error none")

	return
}

func (_PacketConnNone) SetDeadline(_ time.Time) (error error) {
	logger.Warn("this is none of net.PacketConn")

	error = fmt.Errorf("error none")

	return
}

func (_PacketConnNone) SetReadDeadline(_ time.Time) (error error) {
	logger.Warn("this is none of net.PacketConn")

	error = fmt.Errorf("error none")

	return
}

func (_PacketConnNone) SetWriteDeadline(_ time.Time) (error error) {
	logger.Warn("this is none of net.PacketConn")

	error = fmt.Errorf("error none")

	return
}

func (_PacketConnNone) WriteTo(_ []uint8, _ net.Addr) (_ int, error error) {
	logger.Warn("this is none of net.PacketConn")

	error = fmt.Errorf("error none")

	return
}

type PacketConn struct {
	Name string
	Def  net.PacketConn
	Ctx  context.Context
}

func (o *PacketConn) Close() error {
	return PacketConnGet(o.Ctx, o.Name, o.Def).Close()
}

func (o *PacketConn) LocalAddr() net.Addr {
	return PacketConnGet(o.Ctx, o.Name, o.Def).LocalAddr()
}

func (o *PacketConn) ReadFrom(a []uint8) (int, net.Addr, error) {
	return PacketConnGet(o.Ctx, o.Name, o.Def).ReadFrom(a)
}

func (o *PacketConn) SetDeadline(time time.Time) error {
	return PacketConnGet(o.Ctx, o.Name, o.Def).SetDeadline(time)
}

func (o *PacketConn) SetReadDeadline(time time.Time) error {
	return PacketConnGet(o.Ctx, o.Name, o.Def).SetReadDeadline(time)
}

func (o *PacketConn) SetWriteDeadline(time time.Time) error {
	return PacketConnGet(o.Ctx, o.Name, o.Def).SetWriteDeadline(time)
}

func (o *PacketConn) WriteTo(a []uint8, addr net.Addr) (int, error) {
	return PacketConnGet(o.Ctx, o.Name, o.Def).WriteTo(a, addr)
}
