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
	PostInstall(*parser.ConfigurationStruct)
	PreScript(*parser.ConfigurationStruct)
	Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand
	PostScript(*parser.ConfigurationStruct)
	Down(*parser.ConfigurationStruct)
	IsProjectType(*parser.ConfigurationStruct) bool
}

// List is a list of all available plugins
var List = []Plugin{}

// AddPlugin adds a new plugin
func AddPlugin(plugin Plugin) {
	List = append(List, plugin)
}

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
