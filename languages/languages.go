package languages

import (
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
)

type Language interface {
	Setup()
	PreInstall(*parser.ConfigurationStruct)
	Install(*parser.ConfigurationStruct)
	Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand
	IsProjectType(*parser.ConfigurationStruct) bool
}

var LanguageList = []Language{Node{}}

func LanguagesUsed(c *parser.ConfigurationStruct) []Language {
	languagesUsed := []Language{}

	for _, language := range LanguageList {
		if language.IsProjectType(c) {
			languagesUsed = append(languagesUsed, language)
		}
	}

	return languagesUsed
}
