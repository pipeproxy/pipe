package service

import (
	"context"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var service Service
	alias.Register("Service", &service)
}

type Service interface {
	Run(ctx context.Context) error
	Close() error
}
