package load

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/kubernetes-sigs/yaml"
	"github.com/wzshiming/pipe/configure"
)

func Load(ctx context.Context, load io.ReadCloser, i interface{}) error {
	data, err := ioutil.ReadAll(load)
	if err != nil {
		return err
	}
	load.Close()
	data, err = yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return configure.Decode(ctx, data, i)
}
