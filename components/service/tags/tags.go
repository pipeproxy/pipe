package tags

import (
	"context"

	"github.com/pipeproxy/pipe/components/service"
	"github.com/wzshiming/logger"
)

type tags struct {
	service service.Service
	tag     string
	values  []interface{}
}

func newTags(service service.Service, tag string, values []interface{}) *tags {
	return &tags{service: service, tag: tag, values: values}
}

func (t *tags) Run(ctx context.Context) error {
	log := logger.FromContext(ctx)
	if t.tag != "" {
		log.WithName(t.tag)
	}
	if len(t.values) != 0 {
		log.WithValues(t.values)
	}
	ctx = logger.WithContext(ctx, log)
	return t.service.Run(ctx)
}

func (t *tags) Close() error {
	return t.service.Close()
}
