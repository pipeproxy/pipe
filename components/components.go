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
	Decoders         []codec.Decoder
	Encoders         []codec.Encoder
	Marshalers       []codec.Marshaler
	Unmarshalers     []codec.Unmarshaler
	TlsConfigs       []tls.TLS
	Listeners        []listener.ListenConfig
	HttpHandlers     []http.Handler
	ProtocolHandlers []protocol.Handler
	StreamHandlers   []stream.Handler
	Services         []service.Service
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
