package wait

import (
	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/service"
)

const (
	name = "wait"
)

func init() {
	register.Register(name, NewWaitWithConfig)
}

func NewWaitWithConfig() service.Service {
	return newWait()
}
