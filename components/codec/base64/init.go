package base64

import (
	"encoding/base64"

	"github.com/wzshiming/pipe/components/codec"
	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "base64"
)

func init() {
	register.Register(name, NewEncodeWithConfig)
	register.Register(name, NewDecodeWithConfig)
}

type Config struct {
	Encoding string
}

func getEncoding(encoding string) (*base64.Encoding, error) {
	switch encoding {
	case "", "std":
		return base64.StdEncoding, nil
	case "url":
		return base64.URLEncoding, nil
	case "raw_std":
		return base64.RawStdEncoding, nil
	case "raw_url":
		return base64.RawURLEncoding, nil
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
