package base64

import (
	"encoding/base32"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/codec"
)

const (
	name = "base32"
)

func init() {
	decode.Register(name, NewEncodeWithConfig)
	decode.Register(name, NewDecodeWithConfig)
}

type Config struct {
	Encoding string
}

func getEncoding(encoding string) (*base32.Encoding, error) {
	switch encoding {
	case "", "std":
		return base32.StdEncoding, nil
	case "hex":
		return base32.HexEncoding, nil
	default:
		return nil, ErrNotEncoding
	}
}

func NewEncodeWithConfig(conf *Config) (codec.Encoder, error) {
	encoding, err := getEncoding(conf.Encoding)
	if err != nil {
		return nil, err
	}
	return NewCoder(encoding), nil
}

func NewDecodeWithConfig(conf *Config) (codec.Decoder, error) {
	encoding, err := getEncoding(conf.Encoding)
	if err != nil {
		return nil, err
	}
	return NewCoder(encoding), nil
}
