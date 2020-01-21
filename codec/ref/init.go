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
