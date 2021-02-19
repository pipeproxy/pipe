package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/input"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/round_tripper"
)

const (
	name = "http"
)

func init() {
	register.Register(name, NewHTTPWithConfig)
}

type Config struct {
	Dialer stream.Dialer `json:",omitempty"`
	URL    string
}

func NewHTTPWithConfig(conf *Config) input.Input {
	return input.NewLazyReader(func() (input.Input, error) {
		cli := http.Client{
			Transport: round_tripper.RoundTripper(conf.Dialer),
		}
		resp, err := cli.Get(conf.URL)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1024))
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("GET %s fail %s: %s", conf.URL, resp.Status, string(body))
		}
		return input.NewReaderWithAutoClose(resp.Body, resp.Body), nil
	})
}
