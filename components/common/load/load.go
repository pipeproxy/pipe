package load

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/kubernetes-sigs/yaml"
	"github.com/wzshiming/funcfg/kinder"
	"github.com/wzshiming/funcfg/types"
	"github.com/wzshiming/funcfg/unmarshaler"
)

func Load(ctx context.Context, load io.Reader, i interface{}) error {
	data, err := ioutil.ReadAll(load)
	if err != nil {
		return err
	}
	data, err = yaml.YAMLToJSONStrict(data)
	if err != nil {
		return err
	}

	u := unmarshaler.Unmarshaler{
		Ctx:  ctx,
		Get:  types.Get,
		Kind: kinder.Kind,
	}
	return u.Unmarshal(data, i)
}
