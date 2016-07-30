package plugins

import (
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
)

type Plugin interface {
	Setup()
	PreInstall(*parser.ConfigurationStruct)
	Install(*parser.ConfigurationStruct)
	Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand
	IsProjectType(*parser.ConfigurationStruct) bool
}

var PluginList = []Plugin{Node{}}

func PluginsUsed(c *parser.ConfigurationStruct) []Plugin {
	pluginsUsed := []Plugin{}

	for _, plugin := range PluginList {
		if plugin.IsProjectType(c) {
			pluginsUsed = append(pluginsUsed, plugin)
		}
	}

	return pluginsUsed
}
