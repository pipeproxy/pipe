package wait

import (
	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/service"
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
