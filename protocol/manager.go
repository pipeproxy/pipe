package protocol

import (
	"context"
	"fmt"
	"log"
)

var std = newProtocolManager()

func Register(name string, fun NewProtocolFunc) {
	std.Register(name, fun)
}

type handlerManager struct {
	handlers map[string]NewProtocolFunc
}

func newProtocolManager() *handlerManager {
	return &handlerManager{
		handlers: map[string]NewProtocolFunc{},
	}
}

func (h *handlerManager) Register(name string, fun NewProtocolFunc) error {
	log.Printf("[INFO] Register protocol: %s", name)
	h.handlers[name] = fun
	return nil
}

func (h *handlerManager) Get(name string) (NewProtocolFunc, bool) {
	fun, ok := h.handlers[name]
	return fun, ok
}

type NewProtocolFunc func(ctx context.Context, name string, config []byte) (Protocol, error)

func NewProtocol(ctx context.Context, name string, config []byte) (Protocol, error) {
	fun, ok := std.Get(name)
	if !ok {
		return nil, fmt.Errorf("protocol name is not defined: %q: %q", name, config)
	}
	return fun(ctx, name, config)
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
