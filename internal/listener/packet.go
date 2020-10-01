package listener

import (
	"context"
	"net"
	"sync"
	"sync/atomic"

	"github.com/wzshiming/pipe/internal/logger"
)

// ListenPacket returns a net.PacketConn suitable.
func ListenPacket(ctx context.Context, network, address string) (net.PacketConn, error) {
	err := ctx.Err()
	if err != nil {
		return nil, err
	}
	pc, err := listenPacket(ctx, network, address)
	if err != nil {
		return nil, err
	}
	swapClose(ctx, pc)
	return pc, nil
}

func listenPacket(ctx context.Context, network, address string) (*fakeClosePacketConn, error) {
	packetConnMut.Lock()
	defer packetConnMut.Unlock()

	key := buildKey(network, address)
	if _, port, _ := net.SplitHostPort(address); port != "" && port != "0" {
		if global, ok := packetConn[key]; ok {
			atomic.AddInt32(&global.usage, 1)
			logger.Infof("Relisten to %s", key)
			return &fakeClosePacketConn{usage: &global.usage, key: key, PacketConn: global.packetConn}, nil
		}
	}

	var listenConfig net.ListenConfig
	pc, err := listenConfig.ListenPacket(ctx, network, address)
	if err != nil {
		return nil, err
	}
	address = sameAddress(address, pc.LocalAddr().String())
	global := &globalPacketConn{usage: 1, packetConn: pc}
	packetConn[key] = global
	realKey := buildKey(network, address)
	if key != realKey {
		logger.Infof("Listen to %s (%s)", key, realKey)
	} else {
		logger.Infof("Listen to %s", key)
	}
	logger.Infof("Listen to %s", key)
	return &fakeClosePacketConn{usage: &global.usage, key: key, PacketConn: pc}, nil
}

type fakeClosePacketConn struct {
	closed int32
	usage  *int32
	key    string
	net.PacketConn
}

func (f *fakeClosePacketConn) Key() string {
	return f.key
}

func (f *fakeClosePacketConn) Close() error {
	if !atomic.CompareAndSwapInt32(&f.closed, 0, 1) ||
		atomic.AddInt32(f.usage, -1) != 0 {
		return nil
	}

	packetConnMut.Lock()
	defer packetConnMut.Unlock()
	delete(packetConn, f.key)
	logger.Infof("Close listen to %s", f.key)
	return f.PacketConn.Close()
}

type globalPacketConn struct {
	usage      int32
	packetConn net.PacketConn
}

var (
	packetConn    = map[string]*globalPacketConn{}
	packetConnMut sync.Mutex
)
