package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/wzshiming/pipe/init"

	"github.com/wzshiming/pipe/internal/notify"
	"github.com/wzshiming/pipe/internal/stream"

	"github.com/kubernetes-sigs/yaml"
	"github.com/spf13/pflag"
	"github.com/wzshiming/lockfile"
	"github.com/wzshiming/pipe"
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
		log.Printf("[ERROR] lockfile error: %s", err)
		return
	}

	if signal == "" {
		log.Printf("[INFO] Start pipe")
	} else {
		log.Printf("[INFO] Send signal %s to pipe", signal)
	}
	switch signal {
	case "":
		err := lf.Lock()
		if err != nil {
			log.Fatalln("[ERROR] start error:", err)
			return
		}
		start(conf)
		err = lf.Unlock()
		if err != nil {
			log.Fatalln("[ERROR] end error:", err)
			return
		}
	case "reload":
		pid, err := lf.Get()
		if err != nil {
			log.Fatalln("[ERROR] reload error:", err)
			return
		}
		err = notify.Kill(pid, notify.Reload)
		if err != nil {
			log.Fatalln("[ERROR] send error:", err)
			return
		}
	case "stop":
		pid, err := lf.Get()
		if err != nil {
			log.Fatalln("[ERROR] stop error:", err)
			return
		}
		err = notify.Kill(pid, notify.Stop)
		if err != nil {
			log.Fatalln("[ERROR] send error:", err)
			return
		}
	case "reopen":
		pid, err := lf.Get()
		if err != nil {
			log.Fatalln("[ERROR] reopen error:", err)
			return
		}
		err = notify.Kill(pid, notify.Reopen)
		if err != nil {
			log.Fatalln("[ERROR] send error:", err)
			return
		}
	default:
		log.Fatalf("[ERROR] not defined signal %s", signal)
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

	c, err = yaml.YAMLToJSONStrict(c)
	if err != nil {
		return nil, fmt.Errorf("converts config YAML to JSON error: %w", err)
	}
	return c, nil
}

func start(conf string) {

	c, err := getConfig(conf)
	if err != nil {
		log.Printf("[ERROR] read config file %q error: %s", conf, err.Error())
		return
	}

	svc, err := pipe.NewPipeWithConfig(context.Background(), c)
	if err != nil {
		log.Printf("[ERROR] configure config error: %s", err.Error())
		return
	}

	notify.On(notify.Stop, func() {
		log.Println("[INFO] [close] start")
		defer log.Println("[INFO] [close] end")
		defer stream.CloseExcess()

		err := svc.Close()
		if svc == nil {
			log.Printf("[ERROR] service close error: %s", err.Error())
			return
		}
	})
	notify.On(notify.Reload, func() {
		log.Println("[INFO] [reload] start")
		defer log.Println("[INFO] [reload] end")
		defer stream.CloseExcess()

		c, err := getConfig(conf)
		if err != nil {
			log.Printf("[ERROR] read config file %q error: %s", conf, err.Error())
			return
		}

		err = svc.Reload(c)
		if err != nil {
			log.Printf("[ERROR] reload error: %s", err.Error())
			return
		}
	})

	err = svc.Run()
	if err != nil {
		log.Printf("[ERROR] start error: %s", err.Error())
		return
	}

	log.Println("[INFO] exit pipe")
}
