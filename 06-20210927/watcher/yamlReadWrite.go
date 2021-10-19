package watcher

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func DemoYAMLReader1() {
	yfile, err := ioutil.ReadFile("input.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]string)
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func DemoYAMLReader3() {
	yfile, err := ioutil.ReadFile("./hls/result.yaml")
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]VideoNameID)
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		fmt.Printf("%s: %s\n", k, v)
	}
}
