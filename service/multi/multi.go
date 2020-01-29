package multi

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/wzshiming/pipe/service"
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
	case 1:
		return m.multi[0].Run(ctx)
	default:
		m.wg.Add(len(m.multi))
		for _, svc := range m.multi {
			go func(svc service.Service) {
				err := svc.Run(ctx)
				if err != nil {
					log.Printf("[ERROR] service start error: %s", err.Error())
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
	case 1:
		return m.multi[0].Close()
	default:
		for _, service := range m.multi {
			err := service.Close()
			if err != nil {
				log.Printf("[ERROR] service close error: %s", err.Error())
			}
		}
	}
	return nil
}
