package types

import (
	"github.com/pipeproxy/pipe/components/common/load"
	"github.com/wzshiming/logger"
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
			logger.Log.Error(err, "register components")
			return err
		}
	}
	return nil
}
