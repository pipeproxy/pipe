package json

import (
	"encoding/json"
)

type Coder struct {
	buf []byte
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
