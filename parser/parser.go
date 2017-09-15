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
