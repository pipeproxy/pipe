package direct

import (
	"net/http"
	"unsafe"

	"github.com/wzshiming/pipe/configure"
)

const name = "direct"

func init() {
	configure.Register(name, NewDirectWithConfig)
}

type Config struct {
	Code int
	Body string
}

func NewDirectWithConfig(conf *Config) http.Handler {
	return NewDirect(conf.Code, *(*[]byte)(unsafe.Pointer(&conf.Body)))
}
