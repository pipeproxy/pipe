package dialer

import (
	"context"
	"errors"
	"net"

	"github.com/mikioh/tcp"
	"github.com/pipeproxy/pipe/components/balance"
	svc_stream "github.com/pipeproxy/pipe/components/service/stream"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/wzshiming/logger"
)

type Dialer struct {
	network  string
	address  string
	virtual  bool
	original bool
}

func NewDialer(network string, address string, virtual bool, original bool) *Dialer {
	return &Dialer{
		network:  network,
		address:  address,
		virtual:  virtual,
		original: original,
	}
}

func (d *Dialer) DialStream(ctx context.Context) (stream.Stream, error) {
	network := d.network
	address := d.address
	if d.original {
		s, ok := svc_stream.GetRawStreamWithContext(ctx)
		if !ok {
			return nil, errors.New("unable to get raw stream")
		}
		c, err := tcp.NewConn(s)
		if err != nil {
			return nil, err
		}
		addr, err := c.OriginalDst()
		if err == nil {
			address = addr.String()
			h1, port, err := net.SplitHostPort(address)
			if err != nil {
				return nil, err
			}
			localAddr := s.LocalAddr()
			h2, _, err := net.SplitHostPort(localAddr.String())
			if err != nil {
				return nil, err
			}
			if h1 == h2 {
				address = net.JoinHostPort("127.0.0.1", port)
				network = localAddr.Network()
			} else {
				if network == "" {
					network = addr.Network()
				}
			}
		}
	}

	log := logger.FromContext(ctx)
	if d.virtual {
		log.Info("Virtual dial stream",
			"targetAddress", address,
			"virtual", true,
		)
		return listener.VirtualDialContext(ctx, network, address)
	}
	log.Info("Dial stream",
		"targetAddress", address,
	)
	return listener.DialContext(ctx, network, address)
}

func (d *Dialer) Targets() (balance.PolicyEnum, []stream.Dialer) {
	return balance.EnumPolicyNone, []stream.Dialer{d}
}
