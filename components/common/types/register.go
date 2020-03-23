package types

import (
	"log"

	"github.com/wzshiming/pipe/components/common/load"
	"github.com/wzshiming/pipe/components/common/reference"
)

func Register(i interface{}) error {
	list := []func(i interface{}) error{
		load.Register,
		reference.Register,
	}
	for _, item := range list {
		err := item(i)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
