package main

import (
	"flag"

	_ "github.com/pipeproxy/pipe/init"

	"github.com/wzshiming/funcfg/build/bind"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		bind.Bind("")
	} else {
		bind.Bind(args[0])
	}
}
