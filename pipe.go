package pipe

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/once"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/service"
	"golang.org/x/sync/errgroup"
)

type Pipe struct {
	init  []once.Once
	pipe  service.Service
	group errgroup.Group
	ctx   context.Context
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
		init: conf.Init,
		ctx:  ctx,
	}, nil
}

func (c *Pipe) Run() error {
	for _, init := range c.init {
		init.Do()
	}
	c.group.Go(c.pipe.Run)
	return c.group.Wait()
}

func (c *Pipe) Reload(config []byte) error {
	conf := &Config{}
	err := configure.Decode(c.ctx, config, conf)
	if err != nil {
		return err
	}
	if conf.Pipe == nil {
		return fmt.Errorf("no entry pipe field")
	}

	for _, init := range conf.Init {
		init.Do()
	}
	c.group.Go(conf.Pipe.Run)
	c.pipe.Close()
	c.pipe = conf.Pipe
	c.init = conf.Init
	return nil
}

func (c *Pipe) Close() error {
	return c.pipe.Close()
}
