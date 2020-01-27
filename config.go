package pipe

import (
	"github.com/wzshiming/pipe/components"
	"github.com/wzshiming/pipe/once"
	"github.com/wzshiming/pipe/service"
)

type Config struct {
	Pipe       service.Service
	Components components.Components
	Init       []once.Once
}
