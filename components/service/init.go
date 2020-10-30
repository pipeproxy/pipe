package service

import (
	"context"

	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var service Service
	types.Register(&service)
}

type Service interface {
	Run(ctx context.Context) error
	Close() error
}
