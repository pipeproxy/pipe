package pipe

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/service"
	"golang.org/x/sync/errgroup"
)

type Pipe struct {
	pipe  service.Service
	group errgroup.Group
}

func NewPipeWithConfig(ctx context.Context, config []byte) (*Pipe, error) {
	conf := &Config{}
	err := configure.Decode(ctx, config, conf)
	if err != nil {
		return nil, err
	}
	if conf.Pipe == nil {
		return nil, fmt.Errorf("no entry pipe field")
	}

	return &Pipe{
		pipe: conf.Pipe,
	}, nil
}

func (c *Pipe) Run() error {
	c.group.Go(c.pipe.Run)
	return c.group.Wait()
}

func (c *Pipe) Reload(conf []byte) error {
	p, err := NewPipeWithConfig(context.Background(), conf)
	if err != nil {
		return err
	}

	c.group.Go(p.pipe.Run)
	c.pipe.Close()
	c.pipe = p.pipe
	return nil
}

func (c *Pipe) Close() error {
	return c.pipe.Close()
}
