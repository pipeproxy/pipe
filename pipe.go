package pipe

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pipeproxy/pipe/components/common/load"
	"github.com/pipeproxy/pipe/components/once"
	"github.com/pipeproxy/pipe/components/stdio/input/inline"
	"github.com/pipeproxy/pipe/internal/ctxcache"
	"github.com/wzshiming/logger"
	"golang.org/x/sync/errgroup"
	"sigs.k8s.io/yaml"
)

type Pipe struct {
	config        string
	group         *errgroup.Group
	ctx           context.Context
	cancel        []func()
	mut           sync.Mutex
	reloadMut     sync.Mutex
	reloadCounter uint64
	usage         int32
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
	c := &Pipe{
		ctx: ctx,
	}
	c.ctx = context.WithValue(c.ctx, pipeCtxKeyType(0), c)
	c.group, c.ctx = errgroup.WithContext(c.ctx)
	c.config = string(config)
	return c, nil
}

func (c *Pipe) Start() error {
	ctx := c.ctx
	ctx = logger.WithContext(ctx, logger.FromContext(ctx).WithName("first-load"))
	return c.load(ctx, []byte(c.config), true)
}

func (c *Pipe) Run() error {
	err := c.Start()
	if err != nil {
		return err
	}
	return c.Wait()
}

func (c *Pipe) Wait() error {
	return c.group.Wait()
}

func (c *Pipe) waitUsage(i int32) {
	for atomic.LoadInt32(&c.usage) > i {
		time.Sleep(time.Second / 10)
	}
	time.Sleep(time.Second / 2)
}

func (c *Pipe) start(ctx context.Context, o once.Once, first bool) error {
	ctx, cancel := context.WithCancel(ctx)
	c.group.Go(func() error {
		atomic.AddInt32(&c.usage, 1)
		err := o.Do(ctx)
		atomic.AddInt32(&c.usage, -1)
		return err
	})
	if !first {
		c.waitUsage(2)
	}
	c.close()
	c.waitUsage(1)
	c.cancel = append(c.cancel, cancel)
	return nil
}

func (c *Pipe) Reload(config []byte) error {
	ctx := c.ctx
	ctx = logger.WithContext(ctx,
		logger.FromContext(ctx).WithName(
			fmt.Sprintf("reload-%d", atomic.AddUint64(&c.reloadCounter, 1))))
	return c.load(ctx, config, false)
}

func (c *Pipe) load(ctx context.Context, config []byte, first bool) error {
	config, err := yaml.YAMLToJSONStrict(config)
	if err != nil {
		return err
	}
	conf := string(config)
	var o once.Once
	ctx = ctxcache.WithCache(ctx)
	err = load.Load(logger.WithContext(ctx, logger.FromContext(ctx).WithName("loading")),
		inline.NewInlineWithConfig(&inline.Config{Data: conf}), &o)
	if err != nil {
		return err
	}
	if o == nil {
		return fmt.Errorf("no entry")
	}

	c.mut.Lock()
	defer c.mut.Unlock()
	err = c.start(ctx, o, first)
	if err != nil {
		return err
	}
	c.config = conf
	return nil
}

func (c *Pipe) Close() error {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.close()
	c.waitUsage(0)
	return nil
}

func (c *Pipe) close() {
	for _, do := range c.cancel {
		do()
	}
	c.cancel = c.cancel[:0]
	return
}

func (c *Pipe) Config() []byte {
	c.mut.Lock()
	defer c.mut.Unlock()
	return []byte(c.config)
}
