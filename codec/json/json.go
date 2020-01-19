package json

import (
	"context"
	"encoding/json"

	"github.com/wzshiming/pipe/codec"
)

type Coder struct {
	buf []byte
}

func NewEncoder(ctx context.Context, name string, config []byte) (codec.Encoder, error) {
	return NewCoder(nil), nil
}

func NewDecoder(ctx context.Context, name string, config []byte) (codec.Decoder, error) {
	return NewCoder(nil), nil
}

func NewCoder(buf []byte) *Coder {
	return &Coder{buf}
}

func (d *Coder) Decode(v interface{}) error {
	return json.Unmarshal(d.buf, v)
}

func (d *Coder) Encode(v interface{}) error {
	buf, err := json.Marshal(v)
	d.buf = buf
	return err
}

func (d *Coder) Bytes() []byte {
	return d.buf
}
