package pool

import (
	"bytes"
	"sync"
)

const DefaultSize = 32 * 1024

func GetBytes() []byte {
	return pool.GetBytes()
}

func PutBytes(d []byte) {
	pool.PutBytes(d)
}

func GetBuffer() *bytes.Buffer {
	return pool.GetBuffer()
}

func PutBuffer(buf *bytes.Buffer) {
	pool.PutBuffer(buf)
}

var pool = newPools()

type pools struct {
	bytes  sync.Pool
	buffer sync.Pool
}

func newPools() *pools {
	b := &pools{}
	b.bytes.New = func() interface{} {
		return make([]byte, DefaultSize)
	}
	b.buffer.New = func() interface{} {
		return bytes.NewBuffer(b.GetBytes())
	}
	return b
}

func (b *pools) GetBytes() []byte {
	return b.bytes.Get().([]byte)
}

func (b *pools) PutBytes(d []byte) {
	if d == nil || len(d) < DefaultSize {
		return
	}
	d = d[:cap(d)]
	b.bytes.Put(d)
}

func (b *pools) GetBuffer() *bytes.Buffer {
	buf := b.buffer.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (b *pools) PutBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	b.buffer.Put(buf)
}
