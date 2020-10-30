package message

import (
	"context"

	"github.com/pipeproxy/pipe/internal/logger"
)

type Message string

func (m Message) Do(ctx context.Context) error {
	logger.Info(string(m))
	return nil
}
