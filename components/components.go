package components

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"

	"github.com/wzshiming/pipe/codec"
	"github.com/wzshiming/pipe/protocol"
	"github.com/wzshiming/pipe/service"
	"github.com/wzshiming/pipe/stream"
)

type ctxKeyComponents int

type Components struct {
	Protocols      map[string]protocol.Protocol
	Decoders       map[string]codec.Decoder
	Encoders       map[string]codec.Encoder
	StreamHandlers map[string]stream.Handler
	HttpHandlers   map[string]http.Handler
	Listeners      map[string]net.Listener
	Services       map[string]service.Service
	TlsConfigs     map[string]*tls.Config
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
