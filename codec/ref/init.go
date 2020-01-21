package ref

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/codec"

	"github.com/wzshiming/pipe/components"
	"github.com/wzshiming/pipe/configure"
)

var (
	ErrNotCodec = fmt.Errorf("not codec")
)

const name = "ref"

func init() {
	configure.Register(name, NewRefEncoderWithConfig)
	configure.Register(name, NewRefDecoderWithConfig)
	configure.Register(name, NewRefMarshalerWithConfig)
	configure.Register(name, NewRefUnmarshalerWithConfig)
}

type Config struct {
	Ref string
}

func NewRefEncoderWithConfig(ctx context.Context, conf *Config) (codec.Encoder, error) {
	components, ok := components.GetCtxComponents(ctx)
	if !ok || components == nil || components.StreamHandlers == nil {
		return nil, ErrNotCodec
	}
	encoder, ok := components.Encoders[conf.Ref]
	if !ok {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotCodec)
	}
	return encoder, nil
}

func NewRefDecoderWithConfig(ctx context.Context, conf *Config) (codec.Decoder, error) {
	components, ok := components.GetCtxComponents(ctx)
	if !ok || components == nil || components.StreamHandlers == nil {
		return nil, ErrNotCodec
	}
	decoder, ok := components.Decoders[conf.Ref]
	if !ok {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotCodec)
	}
	return decoder, nil
}

func NewRefMarshalerWithConfig(ctx context.Context, conf *Config) (codec.Marshaler, error) {
	components, ok := components.GetCtxComponents(ctx)
	if !ok || components == nil || components.StreamHandlers == nil {
		return nil, ErrNotCodec
	}
	marshaler, ok := components.Marshalers[conf.Ref]
	if !ok {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotCodec)
	}
	return marshaler, nil
}

func NewRefUnmarshalerWithConfig(ctx context.Context, conf *Config) (codec.Unmarshaler, error) {
	components, ok := components.GetCtxComponents(ctx)
	if !ok || components == nil || components.StreamHandlers == nil {
		return nil, ErrNotCodec
	}
	unmarshaler, ok := components.Unmarshalers[conf.Ref]
	if !ok {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotCodec)
	}
	return unmarshaler, nil
}
