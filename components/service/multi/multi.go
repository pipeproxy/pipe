package multi

import (
	"context"
	"fmt"
	"sync"

	"github.com/pipeproxy/pipe/components/service"
	"github.com/wzshiming/logger"
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
		for i, svc := range m.multi {
			go func(ctx context.Context, i int, svc service.Service) {
				log := logger.FromContext(ctx)
				log = log.WithName(getName(i, svc))
				ctx = logger.WithContext(ctx, log)
				err := svc.Run(ctx)
				if err != nil {
					log.Error(err, "service start")
				}
				m.wg.Done()
			}(ctx, i, svc)
		}
		m.wg.Wait()
	}
	return nil
}

func (m *Multi) Close() error {
	switch len(m.multi) {
	case 0:
	default:
		for i, svc := range m.multi {
			err := svc.Close()
			if err != nil {
				logger.Log.WithName(getName(i, svc)).Error(err, "service close")
			}
		}
	}
	return nil
}
