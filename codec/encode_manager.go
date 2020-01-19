package codec

import (
	"context"
	"fmt"
	"log"
)

var stdEncoder = newEncodeManager()

func RegisterEncoder(name string, fun NewEncoderFunc) {
	stdEncoder.Register(name, fun)
}

type handlerManagerEncoder struct {
	handlers map[string]NewEncoderFunc
}

func newEncodeManager() *handlerManagerEncoder {
	return &handlerManagerEncoder{
		handlers: map[string]NewEncoderFunc{},
	}
}

func (h *handlerManagerEncoder) Register(name string, fun NewEncoderFunc) error {
	log.Printf("[INFO] Register codec: encode: %s", name)
	h.handlers[name] = fun
	return nil
}

func (h *handlerManagerEncoder) Get(name string) (NewEncoderFunc, bool) {
	fun, ok := h.handlers[name]
	return fun, ok
}

type NewEncoderFunc func(ctx context.Context, name string, config []byte) (Encoder, error)

func NewEncoder(ctx context.Context, name string, config []byte) (Encoder, error) {
	fun, ok := stdEncoder.Get(name)
	if !ok {
		return nil, fmt.Errorf("encode name is not defined: %q: %q", name, config)
	}
	return fun(ctx, name, config)
}

type Encoder interface {
	Encode(v interface{}) error
	Bytes() []byte
}
