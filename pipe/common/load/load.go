package load

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/kubernetes-sigs/yaml"
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/configure/decode"
)

const name = "load"

func Register(i interface{}) error {
	return decode.RegisterWithBuildFunc(name, Load, i)
}

type Config struct {
	Load io.ReadCloser
}

func Load(ctx context.Context, conf *Config, i interface{}) error {
	data, err := ioutil.ReadAll(conf.Load)
	if err != nil {
		return err
	}
	conf.Load.Close()
	data, err = yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return configure.Decode(ctx, data, i)
}
