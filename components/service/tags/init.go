package tags

import (
	"sort"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/service"
)

const (
	name = "tags"
)

func init() {
	register.Register(name, NewTagsWithConfig)
}

type Config struct {
	Service service.Service
	Tag     string
	Values  map[string]string
}

func NewTagsWithConfig(conf *Config) service.Service {
	var values []interface{}
	if len(conf.Values) != 0 {
		values = make([]interface{}, 0, len(conf.Values)*2)
		keys := make([]string, 0, len(conf.Values))
		for key := range conf.Values {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			values = append(values, key, conf.Values[key])
		}
	}
	return newTags(conf.Service, conf.Tag, values)
}
