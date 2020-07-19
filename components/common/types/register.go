package types

import (
	"github.com/wzshiming/pipe/components/common/load"
	"github.com/wzshiming/pipe/internal/logger"
)

var Global []interface{}

func register(i interface{}) error {
	Global = append(Global, i)
	return nil
}

func Register(i interface{}) error {
	list := []func(i interface{}) error{
		register,
		load.Register,
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
