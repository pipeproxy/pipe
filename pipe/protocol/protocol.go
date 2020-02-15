package protocol

import (
	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var handler Handler
	alias.Register("protocol.Handler", &handler)
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
	ServeProtocol(ptc Protocol)
}
