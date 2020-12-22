package message

import (
	"context"

	"github.com/wzshiming/logger"
)

type Message string

func (m Message) Do(ctx context.Context) error {
	log := logger.FromContext(ctx)
	if log.Enabled() {
		log.Info(string(m))
	}
	return nil
}
