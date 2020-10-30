package protocol

import (
	"context"

	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var handler Handler
	types.Register(&handler)
}

type Protocol interface {
	WriteHeader(Map) error
	WriteBody(interface{}) error
	ReadHeader(Map) error
	ReadBody(interface{}) error
	Close() error
}

type Map interface {
	Add(key string, value string)
	Set(key string, values []string)
	Get(key string) []string
	Del(key string)
	Range(func(key string, values []string) bool)
	Clone() Map
}

type Handler interface {
	ServeProtocol(ctx context.Context, ptc Protocol)
}
