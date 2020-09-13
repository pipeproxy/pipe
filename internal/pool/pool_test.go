package pool

import (
	"testing"
)

func TestBytes(t *testing.T) {
	PutBytes(nil)
	for i := 0; i != 10; i++ {
		buf := GetBytes()
		if len(buf) != DefaultSize || cap(buf) != DefaultSize {
			t.Error(len(buf), cap(buf))
		}
		PutBytes(buf)
	}
}

func TestBuffer(t *testing.T) {
	PutBuffer(nil)
	for i := 0; i != 10; i++ {
		buf := GetBuffer()
		if buf.Len() != 0 || buf.Cap() != DefaultSize {
			t.Error(buf.Len(), buf.Cap())
		}
		PutBuffer(buf)
	}
}
