package types

import (
	"github.com/wzshiming/pipe/components/common/load"
	"github.com/wzshiming/pipe/components/common/reference/def"
	"github.com/wzshiming/pipe/components/common/reference/ref"
	"github.com/wzshiming/pipe/internal/logger"
)

func Register(i interface{}) error {
	list := []func(i interface{}) error{
		load.Register,
		ref.Register,
		def.Register,
	}
	for _, item := range list {
		err := item(i)
		if err != nil {
			logger.Error(err)
			return err
		}
	}
	return nil
}
