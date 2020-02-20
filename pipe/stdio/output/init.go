package output

import (
	"io"

	"github.com/wzshiming/pipe/configure/alias"
	"github.com/wzshiming/pipe/pipe/common/load"
)

func init() {
	var output Output
	alias.Register("Output", &output)
	load.Register(&output)
}

type Output = io.WriteCloser
