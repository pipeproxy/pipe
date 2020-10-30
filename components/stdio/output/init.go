package output

import (
	"io"

	"github.com/pipeproxy/pipe/components/common/types"
)

func init() {
	var output Output
	types.Register(&output)
}

type Output = io.Writer
