package parser

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type version struct {
	Version string
}

type ConfigurationStruct struct {
	Node version
	Go   version
}

func (c *ConfigurationStruct) ParseFile(path string) {

	data, _ := ioutil.ReadFile(path)

	yaml.Unmarshal(data, c)
}
