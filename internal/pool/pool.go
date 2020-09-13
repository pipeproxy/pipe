package pool

import (
	"bytes"
	"sync"
)

const DefaultSize = 32 * 1024

type BytesPool interface {
	Get() []byte
	Put([]byte)
}

type BufferPool interface {
	Get() *bytes.Buffer
	Put(*bytes.Buffer)
}

func GetBytes() []byte {
	return Bytes.Get()
}

func PutBytes(d []byte) {
	Bytes.Put(d)
}

func GetBuffer() *bytes.Buffer {
	return Buffer.Get()
}

func PutBuffer(buf *bytes.Buffer) {
	Buffer.Put(buf)
}

type bytesPool struct {
	sync.Pool
}

func (b *bytesPool) Get() []byte {
	buf := b.Pool.Get().([]byte)
	buf = buf[:cap(buf)]
	return buf
}

func (b *bytesPool) Put(d []byte) {
	if d == nil || len(d) < DefaultSize {
		return
	}
	b.Pool.Put(d)
}

type bufferPool struct {
	sync.Pool
}

func (b *bufferPool) Get() *bytes.Buffer {
	buf := b.Pool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (b *bufferPool) Put(d *bytes.Buffer) {
	if d == nil {
		return
	}
	b.Pool.Put(d)
}

var (
	Bytes BytesPool = &bytesPool{
		Pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, DefaultSize)
			},
		},
	}
	Buffer BufferPool = &bufferPool{
		Pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(Bytes.Get())
			},
		},
	}
)
