package multi

import (
	"github.com/wzshiming/pipe/service"
)

func init() {
	service.Register(name, NewMultiWithConfig)
}
