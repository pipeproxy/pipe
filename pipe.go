package pipe

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/configure"
	"golang.org/x/sync/errgroup"
)

type Pipe struct {
	conf  *Config
	group errgroup.Group
	ctx   context.Context
}

type pipeCtxKeyType int

func GetPipeWithContext(ctx context.Context) (*Pipe, bool) {
	i := ctx.Value(pipeCtxKeyType(0))
	if i == nil {
		return nil, false
	}
	p, ok := i.(*Pipe)
	return p, ok
}

func NewPipeWithConfig(ctx context.Context, config []byte) (*Pipe, error) {
	c := &Pipe{}
	c.conf = &Config{}
	c.ctx = context.WithValue(ctx, pipeCtxKeyType(0), c.conf)
	err := configure.Decode(ctx, config, c.conf)
	if err != nil {
		return nil, err
	}

	if c.conf.Pipe == nil {
		return nil, fmt.Errorf("no entry pipe field")
	}
	return c, nil
}

func (c *Pipe) Run() error {
	c.run()
	return c.group.Wait()
}

func (c *Pipe) run() {
	for _, init := range c.conf.Init {
		init.Do(c.ctx)
	}
	run := func() error {
		return c.conf.Pipe.Run(c.ctx)
	}
	c.group.Go(run)
}

func (c *Pipe) Reload(config []byte) error {
	closeOld := c.conf.Pipe.Close

	conf := &Config{}
	err := configure.Decode(c.ctx, config, conf)
	if err != nil {
		return err
	}
	if conf.Pipe == nil {
		return fmt.Errorf("no entry pipe field")
	}

	defer closeOld()

	c.conf = conf
	c.run()
	return nil
}

func (c *Pipe) Close() error {
	return c.conf.Pipe.Close()
}
