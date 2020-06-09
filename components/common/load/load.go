package load

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/kubernetes-sigs/yaml"
	"github.com/wzshiming/funcfg/kinder"
	"github.com/wzshiming/funcfg/types"
	"github.com/wzshiming/funcfg/unmarshaler"
	"github.com/wzshiming/pipe/components/common/reference/ctxreference"
)

func Load(ctx context.Context, load io.Reader, i interface{}) error {
	ctx = ctxreference.With(ctx)
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
	err = u.Unmarshal(data, i)
	if err != nil {
		return err
	}
	err = ctxreference.Err(ctx)
	if err != nil {
		return err
	}
	return nil
}
