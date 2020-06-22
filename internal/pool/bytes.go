package pool

import (
	"bytes"
	"sync"
)

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

var pool = &bufferPool{
	sync.Pool{
		New: func() interface{} {
			return make([]byte, 64*1024)
		},
	},
}

type bufferPool struct {
	pool sync.Pool
}

func (b *bufferPool) GetBytes() []byte {
	return b.pool.Get().([]byte)
}

func (b *bufferPool) PutBytes(d []byte) {
	if d == nil {
		return
	}
	d = d[:cap(d)]
	b.pool.Put(d)
}

func (b *bufferPool) GetBuffer() *bytes.Buffer {
	buf := bytes.NewBuffer(b.GetBytes())
	buf.Reset()
	return buf
}

func (b *bufferPool) PutBuffer(buf *bytes.Buffer) {
	if buf == nil {
		return
	}
	buf.Reset()
	b.PutBytes(buf.Bytes())
}
