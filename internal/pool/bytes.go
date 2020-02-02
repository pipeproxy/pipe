package pool

import (
	"sync"
)

var Buffer = &bufferPool{
	sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	},
}

type bufferPool struct {
	sync.Pool
}

func (b *bufferPool) Get() []byte {
	return b.Pool.Get().([]byte)
}

func (b *bufferPool) Put(d []byte) {
	b.Pool.Put(d)
}
