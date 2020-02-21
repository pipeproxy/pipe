package pipe

import (
	"context"
	"fmt"
	"sync"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/pipe/once"
	"golang.org/x/sync/errgroup"
)

type Pipe struct {
	config []byte
	group  *errgroup.Group
	ctx    context.Context
	cancel func()
	once   once.Once
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
	c.group, c.ctx = errgroup.WithContext(ctx)
	c.ctx = context.WithValue(c.ctx, pipeCtxKeyType(0), c)
	var o once.Once
	err := configure.Decode(c.ctx, config, &o)
	if err != nil {
		return nil, err
	}
	if o == nil {
		return nil, fmt.Errorf("no entry")
	}
	c.once = o
	c.config = config
	return c, nil
}

func (c *Pipe) Run() error {
	c.mut.Lock()
	err := c.run(c.once)
	if err != nil {
		c.mut.Unlock()
		return err
	}
	c.mut.Unlock()
	return c.group.Wait()
}

func (c *Pipe) run(o once.Once) error {
	ctx, cancel := context.WithCancel(c.ctx)
	run := func() error {
		return o.Do(ctx)
	}
	c.group.Go(run)
	if c.cancel != nil {
		c.cancel()
	}
	c.once = o
	c.cancel = cancel
	return nil
}

func (c *Pipe) Reload(config []byte) error {

	var o once.Once
	err := configure.Decode(c.ctx, config, &o)
	if err != nil {
		return err
	}
	if o == nil {
		return fmt.Errorf("no entry")
	}

	c.mut.Lock()
	defer c.mut.Unlock()
	err = c.run(o)
	if err != nil {
		return err
	}
	c.config = config
	return nil
}

func (c *Pipe) Close() error {
	c.mut.Lock()
	defer c.mut.Unlock()
	if c.cancel != nil {
		c.cancel()
	}
	return nil
}

func (c *Pipe) Config() []byte {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.config
}
