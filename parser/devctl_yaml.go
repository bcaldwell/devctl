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
	Node              Version
	Go                Version
	Scripts           map[string]utilities.RunCommand
	DockerCompose     interface{} `yaml:"docker-compose"`
	DockerComposeFile string      `yaml:"docker-compose-file"`
}

func (c *ConfigurationStruct) ParseFileDefault() error {
	return c.ParseFile("devctl.yaml", "./devctl.yml")
}

func (c *ConfigurationStruct) ParseFile(paths ...string) (err error) {

	for _, path := range paths {
		data, err := ioutil.ReadFile(path)
		if err == nil {
			err = yaml.Unmarshal(data, c)
			return err
		}
	}
	return err
}

func (c *ConfigurationStruct) ParseJson(data string) {
	json.Unmarshal([]byte(data), c)
}
