package listener

import (
	"context"
	"fmt"
	"log"
	"net"
)

var std = newListenerManager()

func Register(name string, fun NewListenerFunc) {
	std.Register(name, fun)
}

type handlerManager struct {
	handlers map[string]NewListenerFunc
}

func newListenerManager() *handlerManager {
	return &handlerManager{
		handlers: map[string]NewListenerFunc{},
	}
}

func (h *handlerManager) Register(name string, fun NewListenerFunc) error {
	log.Printf("[INFO] Register listener: %s", name)
	h.handlers[name] = fun
	return nil
}

func (h *handlerManager) Get(name string) (NewListenerFunc, bool) {
	fun, ok := h.handlers[name]
	return fun, ok
}

type NewListenerFunc func(ctx context.Context, name string, config []byte) (Listener, error)

type Listener = net.Listener

func NewListener(ctx context.Context, name string, config []byte) (Listener, error) {
	fun, ok := std.Get(name)
	if !ok {
		return nil, fmt.Errorf("listener name is not defined: %q: %q", name, config)
	}
	return fun(ctx, name, config)
}
