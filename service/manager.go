package service

import (
	"context"
	"fmt"
	"log"
)

var std = newServerManager()

func Register(name string, fun NewServiceFunc) {
	std.Register(name, fun)
}

type serverManager struct {
	servers map[string]NewServiceFunc
}

func newServerManager() *serverManager {
	return &serverManager{
		servers: map[string]NewServiceFunc{},
	}
}

func (h *serverManager) Register(name string, fun NewServiceFunc) error {
	log.Printf("[INFO] Register service: %s", name)
	h.servers[name] = fun
	return nil
}

func (h *serverManager) Get(name string) (NewServiceFunc, bool) {
	fun, ok := h.servers[name]
	return fun, ok
}

type NewServiceFunc func(ctx context.Context, name string, config []byte) (Service, error)

func NewService(ctx context.Context, name string, config []byte) (Service, error) {
	fun, ok := std.Get(name)
	if !ok {
		return nil, fmt.Errorf("service name is not defined: %q: %q", name, config)
	}
	return fun(ctx, name, config)
}
