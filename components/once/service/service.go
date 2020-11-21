package service

import (
	"context"

	"github.com/pipeproxy/pipe/components/service"
	"github.com/wzshiming/logger"
)

type Service struct {
	svc service.Service
}

func NewService(svc service.Service) *Service {
	return &Service{
		svc: svc,
	}
}

func (m *Service) Do(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		err := m.svc.Close()
		if err != nil {
			logger.FromContext(ctx).Error(err, "service close")
		}
	}()
	return m.svc.Run(ctx)
}
