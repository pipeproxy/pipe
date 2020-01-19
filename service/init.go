package service

import (
	"github.com/wzshiming/pipe/decode"
)

func init() {
	decode.Register(NewService)
}
