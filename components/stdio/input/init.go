package input

import (
	"io"

	"github.com/wzshiming/pipe/components/common/types"
)

func init() {
	var input Input
	types.Register(&input)
}

type Input = io.Reader
