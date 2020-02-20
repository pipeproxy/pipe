package input

import (
	"io"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var input Input
	alias.Register("Input", &input)
	load.Register(&input)
}

type Input = io.ReadCloser
