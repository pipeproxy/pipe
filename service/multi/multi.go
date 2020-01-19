package multi

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/wzshiming/pipe/decode"
	"github.com/wzshiming/pipe/service"
)

var (
	ErrNotServer = fmt.Errorf("not server")
)

type Multi struct {
	wg       sync.WaitGroup
	services []service.Service
}

func NewMultiWithConfig(ctx context.Context, name string, config []byte) (service.Service, error) {
	var conf Config
	err := decode.Decode(ctx, config, &conf)
	if err != nil {
		return nil, err
	}
	switch len(conf.Services) {
	case 1:
		return conf.Services[0], nil
	case 0:
		return nil, ErrNotServer
	}
	return NewMulti(conf.Services)
}

func NewMulti(services []service.Service) (*Multi, error) {
	return &Multi{
		services: services,
	}, nil
}

func (m *Multi) Run() error {
	m.wg.Add(len(m.services))
	for _, svc := range m.services {
		go func(svc service.Service) {
			err := svc.Run()
			if err != nil {
				log.Printf("[ERROR] service start error: %s", err.Error())
			}
			m.wg.Done()
		}(svc)
	}
	m.wg.Wait()
	return nil
}

func (m *Multi) Close() error {
	for _, service := range m.services {
		err := service.Close()
		if err != nil {
			log.Printf("[ERROR] service close error: %s", err.Error())
		}
	}
	return nil
}
