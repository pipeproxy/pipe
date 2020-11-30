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
		for _, svc := range m.multi {
			go func(ctx context.Context, svc service.Service) {
				err := svc.Run(ctx)
				if err != nil {
					logger.FromContext(ctx).Error(err, "service start")
				}
				m.wg.Done()
			}(ctx, svc)
		}
		m.wg.Wait()
	}
	return nil
}

func (m *Multi) Close() error {
	switch len(m.multi) {
	case 0:
	default:
		for _, svc := range m.multi {
			err := svc.Close()
			if err != nil {
				logger.Log.Error(err, "service close")
			}
		}
	}
	return nil
}
