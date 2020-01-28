package pipe

import (
	"github.com/wzshiming/pipe/once"
	"github.com/wzshiming/pipe/service"
)

type Config struct {
	Components []interface{}
	Pipe       service.Service
	Init       []once.Once
}
