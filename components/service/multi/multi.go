package multi

import (
	"context"
	"fmt"
	"sync"

	"github.com/wzshiming/pipe/components/service"
	"github.com/wzshiming/pipe/internal/logger"
)

var (
	ErrNotServer = fmt.Errorf("not server")
)

type Multi struct {
	wg    sync.WaitGroup
	multi []service.Service
}

func NewMulti(multi []service.Service) *Multi {
	return &Multi{
		multi: multi,
	}
}

func (m *Multi) Run(ctx context.Context) error {
	switch len(m.multi) {
	case 0:
	default:
		m.wg.Add(len(m.multi))
		for _, svc := range m.multi {
			go func(svc service.Service) {
				err := svc.Run(ctx)
				if err != nil {
					logger.Errorf("service start error: %s", err.Error())
				}
				m.wg.Done()
			}(svc)
		}
		m.wg.Wait()
	}
	return nil
}

func (m *Multi) Close() error {
	switch len(m.multi) {
	case 0:
	default:
		for _, service := range m.multi {
			err := service.Close()
			if err != nil {
				logger.Errorf("service close error: %s", err.Error())
			}
		}
	}
	return nil
}
