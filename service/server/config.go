package server

import (
	"github.com/wzshiming/pipe/listener"
	"github.com/wzshiming/pipe/stream"
)

type Config struct {
	Listener listener.Listener
	Handlers []stream.Handler
}

const name = "server"
