package output

import (
	"io"

	"github.com/wzshiming/pipe/configure/alias"
)

func init() {
	var output Output
	alias.Register("Output", &output)
}

type Output io.WriteCloser
