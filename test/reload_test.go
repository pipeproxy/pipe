package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/wzshiming/pipe"
	"github.com/wzshiming/pipe/bind"
	"github.com/wzshiming/pipe/internal/stream"
)

func getDirect(port, info string) []byte {
	def := bind.ServiceOnceConfig{
		Service: bind.StreamServiceConfig{
			DisconnectOnClose: true,
			Listener: bind.NetworkStreamListenerListenConfigConfig{
				Network: bind.NetworkStreamListenerListenConfigNetworkEnumEnumTCP,
				Address: port,
			},
			Handler: bind.HTTPStreamHandlerConfig{
				Handler: bind.MuxNetHTTPHandlerConfig{
					Routes: []bind.MuxNetHTTPHandlerRoute{
						{
							Path: "/",
							Handler: bind.DirectNetHTTPHandlerConfig{
								Code: http.StatusOK,
								Body: bind.InlineIoReaderConfig{
									Data: info,
								},
							},
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(def)
	if err != nil {
		panic(err)
	}
	return data
}

func getNone() []byte {
	def := bind.ServiceOnceConfig{
		Service: bind.NoneService{},
	}

	data, err := json.Marshal(def)
	if err != nil {
		panic(err)
	}
	return data
}

func TestReload(t *testing.T) {
	var data = "data"
	p, err := pipe.NewPipeWithConfig(context.Background(), getDirect(":0", data))
	if err != nil {
		t.Fatal(err)
	}
	err = p.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()

	time.Sleep(time.Second / 10)
	d := stream.ListenList()
	if len(d) != 1 {
		t.Fail()
	}

	u, err := url.Parse(d[0])
	if err != nil {
		t.Fatal(err)
	}

	port := u.Port()
	uri := fmt.Sprintf("http://127.0.0.1:%s", port)
	body, err := httpGet(uri)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != data {
		t.FailNow()
	}

	for i := 0; i != 10; i++ {
		data := fmt.Sprintf("data%d", i)
		err := p.Reload(getDirect(fmt.Sprintf(":%s", port), data))
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second / 10)
		stream.CloseExcess()

		body, err := httpGet(uri)
		if err != nil {
			t.Errorf("reload configuration failed times %d, error %s", i, err)
			continue
		}

		if string(body) != data {
			t.Errorf("reload configuration failed times %d, got %q, want %q", i, string(body), data)
			continue
		}

		err = p.Reload(getNone())
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second / 10)
		stream.CloseExcess()

		body, err = httpGet(uri)
		if err == nil && string(body) == data {
			err = fmt.Errorf("port not closed")
			t.Errorf("reload configuration failed times %d, error %s", i, err)
		}
	}
}

func httpGet(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	req = req.WithContext(ctx)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

var httpClient = http.DefaultClient
