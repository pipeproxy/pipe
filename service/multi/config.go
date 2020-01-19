package multi

import (
	"github.com/wzshiming/pipe/service"
)

type Config struct {
	Services []service.Service
}

const name = "multi"
