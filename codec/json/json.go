package json

import (
	"encoding/json"
)

type Coder struct {
}

func NewCoder() Coder {
	return Coder{}
}

func (d Coder) Unmarshal(buf []byte, v interface{}) error {
	return json.Unmarshal(buf, v)
}

func (d Coder) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
