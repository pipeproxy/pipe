package mux

import (
	"github.com/wzshiming/pipe/stream"
)

type Route struct {
	Pattern string
	Regexp  string
	Prefix  string
	Handler stream.Handler
}

type Config struct {
	Routes   []*Route
	NotFound stream.Handler
}

const (
	name = "mux"
)
