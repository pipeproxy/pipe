package input

import (
	"io"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var input Input
	alias.Register("Input", &input)
}

type Input io.ReadCloser
