package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/wzshiming/pipe"
	"github.com/wzshiming/pipe/bind"
)

func getDirect(port, info string) []byte {
	def := bind.ServiceOnceConfig{
		Service: bind.StreamServiceConfig{
			DisconnectOnClose: true,
			Listener: bind.ListenerStreamListenConfigConfig{
				Network: bind.ListenerStreamListenConfigListenerNetworkEnumEnumTCP,
				Address: port,
			},
			Handler: bind.HTTP1StreamHandlerConfig{
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

func getWait() []byte {
	def := bind.ServiceOnceConfig{
		Service: bind.WaitService{},
	}

	data, err := json.Marshal(def)
	if err != nil {
		panic(err)
	}
	return data
}

func TestReload(t *testing.T) {
	var port = "8888"
	var data = "data"
	svc, err := pipe.NewPipeWithConfig(context.Background(), getDirect(fmt.Sprintf(":%s", port), data))
	if err != nil {
		t.Fatal(err)
	}

	err = svc.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer svc.Close()
	time.Sleep(time.Second / 5)

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
		err := svc.Reload(getDirect(fmt.Sprintf(":%s", port), data))
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)

		body, err := httpGet(uri)
		if err != nil {
			t.Errorf("reload configuration failed times %d, error %s", i, err)
			continue
		}

		if string(body) != data {
			t.Errorf("reload configuration failed times %d, got %q, want %q", i, string(body), data)
			continue
		}

		err = svc.Reload(getWait())
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)

		body, err = httpGet(uri)
		if err == nil && string(body) == data {
			err = fmt.Errorf("port not closed")
			t.Errorf("reload configuration failed times %d, error %s", i, err)
		}

		err = svc.Reload(getDirect(":0", data))
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)

		body, err = httpGet(uri)
		if err == nil && string(body) == data {
			err = fmt.Errorf("port not closed")
			t.Errorf("reload configuration failed times %d, error %s", i, err)
		}

		err = svc.Reload(getDirect(fmt.Sprintf(":%s", port), data))
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second)

		body, err = httpGet(uri)
		if err != nil {
			t.Errorf("reload configuration failed times %d, error %s", i, err)
			continue
		}

		if string(body) != data {
			t.Errorf("reload configuration failed times %d, got %q, want %q", i, string(body), data)
			continue
		}
	}
}

func httpGet(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithDeadline(req.Context(), time.Now().Add(time.Second/5))
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
