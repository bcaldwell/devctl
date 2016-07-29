package parser

import (
	"io/ioutil"

	"github.com/benjamincaldwell/devctl/utilities"

	"gopkg.in/yaml.v2"
)

type version struct {
	Version string
}

type ConfigurationStruct struct {
	Node    version
	Go      version
	Scripts map[string]utilities.RunCommand
}

func (c *ConfigurationStruct) ParseFile(path string) {

	data, _ := ioutil.ReadFile(path)

	yaml.Unmarshal(data, c)
}
