package parser

import (
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

var DevctlConfig *Config

type Config struct {
	GithubUser string `json:"github_user"`
	SourceDir  string `json:"source_dir"`
	GitlabUser string `json:"gitlab_user"`
	GitlabURL  string `json:"gitlab_url"`
}

func ReadInConfig(paths ...string) (err error) {
	DevctlConfig = new(Config)
	for _, path := range paths {
		data, err := ioutil.ReadFile(path)
		data_s := os.ExpandEnv(string(data))

		if err == nil {
			err = yaml.Unmarshal([]byte(data_s), DevctlConfig)
			return err
		}
	}
	return err
}
