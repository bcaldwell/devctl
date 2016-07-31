package parser

import (
	"encoding/json"
	"io/ioutil"

	"github.com/benjamincaldwell/devctl/utilities"

	"gopkg.in/yaml.v2"
)

type Version struct {
	Version string
}

type ConfigurationStruct struct {
	Node    Version
	Go      Version
	Scripts map[string]utilities.RunCommand
}

func (c *ConfigurationStruct) ParseFile(path string) {

	data, _ := ioutil.ReadFile(path)

	yaml.Unmarshal(data, c)
}

func (c *ConfigurationStruct) ParseJson(data string) {
	json.Unmarshal([]byte(data), c)
}
