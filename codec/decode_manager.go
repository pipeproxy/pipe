package codec

import (
	"context"
	"fmt"
	"log"
)

var stdDecoder = newDecodeManager()

func RegisterDecoder(name string, fun NewDecoderFunc) {
	stdDecoder.Register(name, fun)
}

type handlerManagerDecoder struct {
	handlers map[string]NewDecoderFunc
}

func newDecodeManager() *handlerManagerDecoder {
	return &handlerManagerDecoder{
		handlers: map[string]NewDecoderFunc{},
	}
}

func (h *handlerManagerDecoder) Register(name string, fun NewDecoderFunc) error {
	log.Printf("[INFO] Register codec: decode: %s", name)
	h.handlers[name] = fun
	return nil
}

func (h *handlerManagerDecoder) Get(name string) (NewDecoderFunc, bool) {
	fun, ok := h.handlers[name]
	return fun, ok
}

type NewDecoderFunc func(ctx context.Context, name string, config []byte) (Decoder, error)

func NewDecoder(ctx context.Context, name string, config []byte) (Decoder, error) {
	fun, ok := stdDecoder.Get(name)
	if !ok {
		return nil, fmt.Errorf("decode name is not defined: %q: %q", name, config)
	}
	return fun(ctx, name, config)
}

type Decoder interface {
	Decode(v interface{}) error
	Bytes() []byte
}
