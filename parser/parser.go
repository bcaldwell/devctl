package parser

import (
	"io/ioutil"

	"github.com/bitly/go-simplejson"
)

// https://github.com/buger/jsonparser
// https://github.com/tidwall/gjson

var configData map[string]string

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

// ParseDevctlConfig parses main devctl config file. Data is stored in parser package for future access with GetString method
func ParseDevctlConfig(file string) (err error) {
	configData, err = ReadTomlLike(file)
	return err
}

// GetString return the key from the config as a string. Optional default value as second argument.
func GetString(key string, params ...string) string {
	value, ok := configData[key]
	if !ok {
		if len(params) > 0 {
			value = params[0]
		} else {
			value = ""
		}
	}
	return value
}
