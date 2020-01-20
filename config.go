package pipe

import (
	"github.com/wzshiming/pipe/components"
	"github.com/wzshiming/pipe/service"
)

type Config struct {
	Pipe service.Service
}

type configComponents struct {
	Components *components.Components
}
