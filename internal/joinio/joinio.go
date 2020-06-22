package joinio

import (
	"io"

	"github.com/wzshiming/pipe/internal/pool"
)

func BothCopy(p1, p2 io.ReadWriter) error {
	go Copy(p1, p2)
	Copy(p2, p1)
	return nil
}

func Copy(p1 io.Writer, p2 io.Reader) error {
	buf := pool.GetBytes()
	io.CopyBuffer(p1, p2, buf[:])
	pool.PutBytes(buf)
	return nil
}
