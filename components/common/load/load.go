package load

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/wzshiming/funcfg/types"
	"github.com/wzshiming/funcfg/unmarshaler"
	"github.com/wzshiming/pipe/internal/refctx"
	"sigs.k8s.io/yaml"
)

func Load(ctx context.Context, load io.Reader, i interface{}) error {
	ctx = refctx.With(ctx)
	data, err := ioutil.ReadAll(load)
	if err != nil {
		return err
	}
	data, err = yaml.YAMLToJSONStrict(data)
	if err != nil {
		return err
	}
	u := unmarshaler.Unmarshaler{
		Ctx:      ctx,
		Provider: types.Default,
	}
	err = u.Unmarshal(data, i)
	if err != nil {
		return err
	}
	err = refctx.Err(ctx)
	if err != nil {
		return err
	}
	return nil
}
