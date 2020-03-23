package message

import (
	"context"
	"log"
)

type Message string

func (m Message) Do(ctx context.Context) error {
	log.Println(string(m))
	return nil
}
