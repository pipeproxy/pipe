package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/pipeproxy/pipe/init"
	_ "github.com/pipeproxy/pipe/internal/log"

	"github.com/pipeproxy/pipe"
	"github.com/pipeproxy/pipe/internal/notify"
	"github.com/spf13/pflag"
	"github.com/wzshiming/lockfile"
	"github.com/wzshiming/logger"
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

	log := logger.Log
	lf, err := lockfile.NewLockfile(pidfile)
	if err != nil {
		log.Error(err, "lockfile")
		return
	}

	if signal == "" {
		log.Info("Start pipe")
	} else {
		log = log.WithName(signal).WithValues("signal", signal)
		log.Info("send signal to pipe")
	}
	switch signal {
	case "":
		err := lf.Lock()
		if err != nil {
			log.Error(err, "lock pidfile")
			return
		}
		ctx := logger.WithContext(context.Background(), log)
		start(ctx, log, conf)
		err = lf.Unlock()
		if err != nil {
			log.Error(err, "unlock pidfile")
			return
		}
	case "reload":
		pid, err := lf.Get()
		if err != nil {
			log.Error(err, "reload")
			return
		}
		err = notify.Kill(pid, notify.Reload)
		if err != nil {
			log.Error(err, "send signal reload")
			return
		}
	case "stop":
		pid, err := lf.Get()
		if err != nil {
			log.Error(err, "stop")
			return
		}
		err = notify.Kill(pid, notify.Stop)
		if err != nil {
			log.Error(err, "send signal stop")
			return
		}
	case "reopen":
		pid, err := lf.Get()
		if err != nil {
			log.Error(err, "reopen")
			return
		}
		err = notify.Kill(pid, notify.Reopen)
		if err != nil {
			log.Error(err, "send signal reopen")
			return
		}
	default:
		log.V(-2).Info("not defined signal")
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

func start(ctx context.Context, log logger.Logger, conf string) {

	c, err := getConfig(conf)
	if err != nil {
		log.Error(err, "read config file",
			"config", conf,
		)
		return
	}

	svc, err := pipe.NewPipeWithConfig(ctx, c)
	if err != nil {
		log.Error(err, "configure config")
		return
	}

	notify.On(notify.Stop, func() {
		log.Info("Closing")

		err := svc.Close()
		if err != nil {
			log.Error(err, "service close")
			return
		}
	})
	notify.On(notify.Reload, func() {
		log.Info("Reloading")

		c, err := getConfig(conf)
		if err != nil {
			log.Error(err, "read config file",
				"config", conf,
			)
			return
		}

		err = svc.Reload(c)
		if err != nil {
			log.Error(err, "reload")
			return
		}
	})

	err = svc.Run()
	if err != nil {
		log.Error(err, "start")
		return
	}

	log.Info("Exit pipe")
}
