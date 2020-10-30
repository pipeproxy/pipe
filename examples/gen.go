//+build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pipeproxy/pipe/examples"
	"sigs.k8s.io/yaml"
)

func main() {
	for file, v := range examples.Examples {
		d, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		d, err = yaml.JSONToYAML(d)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ioutil.WriteFile(file+".yml", d, 0655)
	}
}
