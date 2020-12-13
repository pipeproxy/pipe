package stream

import (
	"context"
	"net"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var handler Handler
	types.Register(&handler)
	var listenConfig ListenConfig
	types.Register(&listenConfig)
	var dialer Dialer
	types.Register(&dialer)
	var stream Stream
	types.Register(&stream)
}

type Stream = net.Conn

type StreamListener = net.Listener

type Handler interface {
	ServeStream(ctx context.Context, stm Stream)
}

type ListenConfig interface {
	ListenStream(ctx context.Context) (StreamListener, error)
	IsVirtual() bool
}

type Dialer interface {
	DialStream(ctx context.Context) (Stream, error)
	IsVirtual() bool
	Targets() []Dialer
	Policy() balance.Policy
	String() string
}

type NetworkEnum string

const (
	EnumNetworkTCP  NetworkEnum = "tcp"
	EnumNetworkTCP4 NetworkEnum = "tcp4"
	EnumNetworkTCP6 NetworkEnum = "tcp6"
	EnumNetworkUnix NetworkEnum = "unix"
)
