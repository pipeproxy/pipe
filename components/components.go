package components

import (
	"context"
	"net/http"

	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/listener"
	"github.com/wzshiming/pipe/protocol"
	"github.com/wzshiming/pipe/service"
	"github.com/wzshiming/pipe/stream"
	"github.com/wzshiming/pipe/tls"
)

type ctxKeyComponents int

type Components struct {
	Decoders         map[string]codec.Decoder
	Encoders         map[string]codec.Encoder
	Marshalers       map[string]codec.Marshaler
	Unmarshalers     map[string]codec.Unmarshaler
	Listeners        map[string]listener.ListenConfig
	Services         map[string]service.Service
	TlsConfigs       map[string]tls.TLS
	StreamHandlers   map[string]stream.Handler
	HttpHandlers     map[string]http.Handler
	ProtocolHandlers map[string]protocol.Handler
}

func PutCtxComponents(ctx context.Context, components *Components) context.Context {
	return context.WithValue(ctx, ctxKeyComponents(0), components)
}

func GetCtxComponents(ctx context.Context) (components *Components, ok bool) {
	d := ctx.Value(ctxKeyComponents(0))
	if d == nil {
		return nil, false
	}
	components, ok = d.(*Components)
	if !ok || components == nil {
		return nil, false
	}
	return components, true
}
