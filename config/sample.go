package config

import (
	"github.com/wzshiming/pipe/bind"
)

func BuildSampleWithOnce(start bind.Service, components ...bind.Component) bind.Once {
	if len(components) == 0 {
		return bind.ServiceOnceConfig{
			Service: start,
		}
	}
	multi := []bind.Once{
		bind.ComponentsOnceConfig{
			Components: components,
		},
		bind.ServiceOnceConfig{
			Service: start,
		},
	}
	return bind.MultiOnceConfig{
		Multi: multi,
	}
}
