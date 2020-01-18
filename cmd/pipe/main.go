package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/wzshiming/pipe/init"

	"github.com/kubernetes-sigs/yaml"
	"github.com/spf13/pflag"
	"github.com/wzshiming/pipe"
	"github.com/wzshiming/pipe/decode"
)

var conf string

func init() {
	pipeConfig := "./pipe.yml"
	if n, ok := os.LookupEnv("PIPE_CONFIG"); ok {
		pipeConfig = n
	}

	pflag.StringVarP(&conf, "config", "c", pipeConfig, "Use an alternative configuration file.")

	pflag.Parse()
}

func main() {

	c, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Printf("[ERROR] reader config file error: %s", err.Error())
		return
	}

	c, err = yaml.YAMLToJSON(c)
	if err != nil {
		log.Printf("[ERROR] converts config YAML to JSON error: %s", err.Error())
		return
	}

	ctx := context.Background()

	var conf pipe.Config
	err = decode.Decode(ctx, c, &conf)
	if err != nil {
		log.Printf("[ERROR] decode config error: %s", err.Error())
		return
	}

	for _, server := range conf.Servers {
		err := server.Start()
		if err != nil {
			log.Printf("[ERROR] start error: %s", err.Error())
			return
		}
	}
	<-make(chan struct{})
}
