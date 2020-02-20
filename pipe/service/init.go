package service

import (
	"context"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var service Service
	alias.Register("Service", &service)
	load.Register(&service)
}

type Service interface {
	Run(ctx context.Context) error
	Close() error
}
