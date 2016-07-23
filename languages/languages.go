package languages

import "github.com/benjamincaldwell/devctl/parser"

type Language interface {
	Setup()
	PreInstall(*parser.ConfigurationStruct)
	Install(*parser.ConfigurationStruct)
}

var LanguageList = []Language{Node{}}
