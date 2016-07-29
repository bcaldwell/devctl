package parser

import (
	"io/ioutil"

	"github.com/bitly/go-simplejson"
	"github.com/spf13/viper"
)

type Parser struct {
	Data *simplejson.Json
}

func (p *Parser) Parse(file string) error {
	data, err := ioutil.ReadFile(file)
	if err == nil {
		p.Data, _ = simplejson.NewJson(data)
		return nil
	}
	return err
}

func (p *Parser) GetMap(key string) map[string]interface{} {
	return p.Data.Get(key).MustMap()
}

// GetString return the key from the config as a string. Optional default value as second argument.
func GetString(key string, params ...string) (value string) {
	value = viper.GetString(key)
	if value == "" {
		value = params[0]
	}
	return
}
