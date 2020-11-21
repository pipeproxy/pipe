package message

import (
	"context"

	"github.com/wzshiming/logger"
)

type Message string

func (m Message) Do(ctx context.Context) error {
	logger.FromContext(ctx).Info(string(m))
	return nil
}
