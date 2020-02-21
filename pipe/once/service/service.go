package service

import (
	"context"

	"github.com/wzshiming/pipe/pipe/service"
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
	return m.svc.Run(ctx)
}
