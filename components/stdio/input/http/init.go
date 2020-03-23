package http

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/protocol/http/round_tripper"
	"github.com/wzshiming/pipe/components/stdio/input"
)

const name = "http"

func init() {
	register.Register(name, NewHTTPWithConfig)
}

type Config struct {
	RoundTripper round_tripper.RoundTripper
	URL          string
}

var defaultTransport = http.DefaultTransport.(*http.Transport)

func NewHTTPWithConfig(conf *Config) (input.Input, error) {
	roundTripper := conf.RoundTripper
	if roundTripper == nil {
		roundTripper = http.DefaultTransport
	}
	cli := http.Client{
		Transport: roundTripper,
	}
	resp, err := cli.Get(conf.URL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return ioutil.NopCloser(bytes.NewReader(body)), nil
}
