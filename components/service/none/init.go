package none

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/service"
)

const (
	name = "none"
)

func init() {
	register.Register(name, NewNoneWithConfig)
}

func NewNoneWithConfig() service.Service {
	return newNone()
}
