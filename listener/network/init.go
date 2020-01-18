package network

import (
	"github.com/wzshiming/pipe/listener"
)

func init() {
	listener.Register("network", NewNetwork)
}
