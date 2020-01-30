package pipe

import (
	"context"
	"fmt"
	"sync"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/once"
	"github.com/wzshiming/pipe/service"
	"golang.org/x/sync/errgroup"
)

type Pipe struct {
	conf  *Config
	group errgroup.Group
	ctx   context.Context
	pipe  service.Service
	init  []once.Once
	mut   sync.Mutex
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
	c.ctx = context.WithValue(ctx, pipeCtxKeyType(0), c)
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
	err := c.run(c.conf.Pipe, c.conf.Init)
	if err != nil {
		return err
	}
	return c.group.Wait()
}

func (c *Pipe) run(pipe service.Service, init []once.Once) error {
	c.mut.Lock()
	defer c.mut.Unlock()

	for _, init := range init {
		err := init.Do(c.ctx)
		if err != nil {
			return err
		}
	}
	run := func() error {
		return pipe.Run(c.ctx)
	}
	c.group.Go(run)

	if c.pipe != nil {
		c := c.pipe.Close
		defer c()
	}
	c.init = init
	c.pipe = pipe
	return nil
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

	err = c.run(conf.Pipe, conf.Init)
	if err != nil {
		return err
	}

	c.mut.Lock()
	defer c.mut.Unlock()
	c.conf = conf
	return nil
}

func (c *Pipe) Close() error {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.pipe.Close()
}
