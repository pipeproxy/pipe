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
	conf   *Config
	config []byte
	group  *errgroup.Group
	ctx    context.Context
	cancel func()
	pipe   service.Service
	init   []once.Once
	mut    sync.Mutex
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
	c.config = config
	c.group, c.ctx = errgroup.WithContext(ctx)
	c.ctx = context.WithValue(c.ctx, pipeCtxKeyType(0), c)
	err := configure.Decode(c.ctx, config, c.conf)
	if err != nil {
		return nil, err
	}

	if c.conf.Pipe == nil {
		return nil, fmt.Errorf("no entry pipe field")
	}
	return c, nil
}

func (c *Pipe) Run() error {
	c.mut.Lock()
	err := c.run(c.conf.Pipe, c.conf.Init)
	if err != nil {
		c.mut.Unlock()
		return err
	}
	c.mut.Unlock()
	return c.group.Wait()
}

func (c *Pipe) run(pipe service.Service, init []once.Once) error {
	ctx, cancel := context.WithCancel(c.ctx)

	for _, init := range init {
		err := init.Do(ctx)
		if err != nil {
			return err
		}
	}
	run := func() error {
		return pipe.Run(ctx)
	}
	c.group.Go(run)

	if c.cancel != nil {
		c.cancel()
	}
	if c.pipe != nil {
		c.pipe.Close()
	}

	c.init = init
	c.pipe = pipe
	c.cancel = cancel
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

	c.mut.Lock()
	defer c.mut.Unlock()
	err = c.run(conf.Pipe, conf.Init)
	if err != nil {
		return err
	}
	c.conf = conf
	c.config = config
	return nil
}

func (c *Pipe) Close() error {
	c.mut.Lock()
	defer c.mut.Unlock()
	if c.cancel != nil {
		c.cancel()
	}
	return c.pipe.Close()
}

func (c *Pipe) Config() []byte {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.config
}
