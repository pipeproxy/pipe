package pipe

import (
	"github.com/wzshiming/pipe/pipe/once"
	"github.com/wzshiming/pipe/pipe/service"
)

type Config struct {
	Components []interface{}
	Pipe       service.Service
	Init       []once.Once
}
