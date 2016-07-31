package plugins

import (
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
)

// Plugin interfaceUse
type Plugin interface {
	Setup()
	PreInstall(*parser.ConfigurationStruct)
	Install(*parser.ConfigurationStruct)
	Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand
	IsProjectType(*parser.ConfigurationStruct) bool
}

// List is a list of all available plugins
var List = []Plugin{Node{}}

// Used calls IsProjectType on each element on the PluginList array and returns a filtered version
func Used(c *parser.ConfigurationStruct) []Plugin {
	pluginsUsed := []Plugin{}

	for _, plugin := range List {
		if plugin.IsProjectType(c) {
			pluginsUsed = append(pluginsUsed, plugin)
		}
	}
	return pluginsUsed
}
