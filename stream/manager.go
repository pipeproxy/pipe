package stream

import (
	"context"
	"fmt"
	"log"
)

var std = newHandlerManager()

func Register(name string, fun NewHandlerFunc) {
	std.Register(name, fun)
}

type handlerManager struct {
	handlers map[string]NewHandlerFunc
}

func newHandlerManager() *handlerManager {
	return &handlerManager{
		handlers: map[string]NewHandlerFunc{},
	}
}

func (h *handlerManager) Register(name string, fun NewHandlerFunc) error {
	log.Printf("[INFO] Register stream: %s", name)
	h.handlers[name] = fun
	return nil
}

func (h *handlerManager) Get(name string) (NewHandlerFunc, bool) {
	fun, ok := h.handlers[name]
	return fun, ok
}

type NewHandlerFunc func(ctx context.Context, name string, config []byte) (Handler, error)

func NewHandler(ctx context.Context, name string, config []byte) (Handler, error) {
	fun, ok := std.Get(name)
	if !ok {
		return nil, fmt.Errorf("handler name is not defined: %q: %q", name, config)
	}
	return fun(ctx, name, config)
}
