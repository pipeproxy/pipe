package tunnel

import (
	"context"
	"io"

	"github.com/wzshiming/pipe/internal/pool"
)

func Tunnel(ctx context.Context, c1, c2 io.ReadWriter) error {
	ctx, cancel := context.WithCancel(ctx)
	var errs tunnelErr
	go func() {
		buf := pool.GetBytes()
		defer pool.PutBytes(buf)
		_, errs[0] = io.CopyBuffer(c1, c2, buf)
		cancel()
	}()
	go func() {
		buf := pool.GetBytes()
		defer pool.PutBytes(buf)
		_, errs[1] = io.CopyBuffer(c2, c1, buf)
		cancel()
	}()
	<-ctx.Done()
	errs[2] = ctx.Err()
	if errs[2] == context.Canceled {
		errs[2] = nil
	}
	return errs.FirstError()
}

func TunnelWithClose(ctx context.Context, c1, c2 io.ReadWriteCloser) error {
	var errs tunnelErr
	errs[0] = Tunnel(ctx, c1, c2)
	errs[1] = c1.Close()
	errs[2] = c2.Close()
	return errs.FirstError()
}

type tunnelErr [3]error

func (t tunnelErr) FirstError() error {
	for _, err := range t {
		if err != nil {
			return err
		}
	}
	return nil
}
