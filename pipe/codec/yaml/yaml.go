package yaml

import (
	"github.com/kubernetes-sigs/yaml"
)

type Coder struct {
}

func NewCoder() Coder {
	return Coder{}
}

func (d Coder) Unmarshal(buf []byte, v interface{}) error {
	return yaml.UnmarshalStrict(buf, v)
}

func (d Coder) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
