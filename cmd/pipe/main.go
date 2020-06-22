package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/wzshiming/pipe/init"

	"github.com/spf13/pflag"
	"github.com/wzshiming/lockfile"
	"github.com/wzshiming/pipe"
	"github.com/wzshiming/pipe/internal/logger"
	"github.com/wzshiming/pipe/internal/notify"
	"github.com/wzshiming/pipe/internal/stream"
)

var signal string
var conf string
var pidfile string

func init() {
	pipeConfig := "./pipe.yml"
	if n, ok := os.LookupEnv("PIPE_CONFIG"); ok {
		pipeConfig = n
	}
	pipePid := "./pipe.pid"
	if n, ok := os.LookupEnv("PIPE_PIDFILE"); ok {
		pipePid = n
	}

	pflag.StringVarP(&signal, "signal", "s", "", "reopen|reload|stop Send a signal to the master process.")
	pflag.StringVarP(&conf, "config", "c", pipeConfig, "Use an alternative configuration file.")
	pflag.StringVarP(&pidfile, "pidfile", "p", pipePid, "contains the process ID of pipe.")

	pflag.Parse()
}

func main() {

	lf, err := lockfile.NewLockfile(pidfile)
	if err != nil {
		logger.Fatalf("lockfile error: %s", err)
		return
	}

	if signal == "" {
		logger.Infof("Start pipe")
	} else {
		logger.Infof("Send signal %s to pipe", signal)
	}
	switch signal {
	case "":
		err := lf.Lock()
		if err != nil {
			logger.Fatal("start error:", err)
			return
		}
		start(conf)
		err = lf.Unlock()
		if err != nil {
			logger.Fatal("end error:", err)
			return
		}
	case "reload":
		pid, err := lf.Get()
		if err != nil {
			logger.Fatal("reload error:", err)
			return
		}
		err = notify.Kill(pid, notify.Reload)
		if err != nil {
			logger.Fatal("send error:", err)
			return
		}
	case "stop":
		pid, err := lf.Get()
		if err != nil {
			logger.Fatal("stop error:", err)
			return
		}
		err = notify.Kill(pid, notify.Stop)
		if err != nil {
			logger.Fatal("send error:", err)
			return
		}
	case "reopen":
		pid, err := lf.Get()
		if err != nil {
			logger.Fatal("reopen error:", err)
			return
		}
		err = notify.Kill(pid, notify.Reopen)
		if err != nil {
			logger.Fatal("send error:", err)
			return
		}
	default:
		logger.Fatalf("not defined signal %s", signal)
		return
	}
}

func getConfig(conf string) ([]byte, error) {
	conf, err := filepath.Abs(conf)
	if err != nil {
		return nil, err
	}

	c, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, fmt.Errorf("read config file %q error: %w", conf, err)
	}

	return c, nil
}

func start(conf string) {

	c, err := getConfig(conf)
	if err != nil {
		logger.Errorf("read config file %q error: %s", conf, err)
		return
	}

	svc, err := pipe.NewPipeWithConfig(context.Background(), c)
	if err != nil {
		logger.Errorf("configure config error: %s", err)
		return
	}

	notify.On(notify.Stop, func() {
		logger.Info("[close] start")
		defer logger.Info("[close] end")
		defer stream.CloseExcess()

		err := svc.Close()
		if svc == nil {
			logger.Errorf("service close error: %s", err)
			return
		}
	})
	notify.On(notify.Reload, func() {
		logger.Info("[reload] start")
		defer logger.Info("[reload] end")
		defer stream.CloseExcess()

		c, err := getConfig(conf)
		if err != nil {
			logger.Errorf("read config file %q error: %s", conf, err)
			return
		}

		err = svc.Reload(c)
		if err != nil {
			logger.Errorf("reload error: %s", err)
			return
		}
	})

	err = svc.Run()
	if err != nil {
		logger.Errorf("start error: %s", err)
		return
	}

	logger.Info("exit pipe")
}
