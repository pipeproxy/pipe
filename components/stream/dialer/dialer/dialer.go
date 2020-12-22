package dialer

import (
	"context"
	"net"

	"github.com/pipeproxy/pipe/components/balance"
	"github.com/pipeproxy/pipe/components/balance/none"
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
	name     string
}

func NewDialer(network string, address string, virtual bool, original bool) *Dialer {
	d := &Dialer{
		network:  network,
		address:  address,
		virtual:  virtual,
		original: original,
	}
	d.name = d.getName()
	return d
}

func (d *Dialer) DialStream(ctx context.Context) (stream.Stream, error) {
	network := d.network
	address := d.address
	if d.original {
		s, addr, ok := svc_stream.GetRawStreamAndOriginalDestinationAddrWithContext(ctx)
		if ok {
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
	if log.Enabled() {
		if d.virtual {
			log.Info("Virtual dial stream",
				"targetAddress", address,
				"virtual", true,
			)
		} else {
			log.Info("Dial stream",
				"targetAddress", address,
			)
		}

	}
	if d.virtual {
		return listener.VirtualDialContext(ctx, network, address)
	}
	return listener.DialContext(ctx, network, address)
}

func (d *Dialer) IsVirtual() bool {
	return d.virtual
}

func (d *Dialer) Targets() []stream.Dialer {
	return []stream.Dialer{d}
}

func (d *Dialer) Policy() balance.Policy {
	return none.NewNone()
}

func (d *Dialer) String() string {
	return d.name
}

func (d *Dialer) getName() string {
	name := d.network + "://" + d.address
	if d.virtual || d.original {
		name += "?"
		if d.virtual {
			name += "virtual"
		}
		if d.original {
			if d.virtual {
				name += "&"
			}
			name = "original"
		}
	}
	return name
}
