package plugins

import (
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/utilities"
)

// Plugin interfaceUse
type Plugin interface {
	// implement stringer interface
	String() string

	// Runs on devctl install or upgrade
	Setup()

	// Tasks initializes tasks required during the up stage.
	// The tasks returned by Tasks will be run in order
	// Note: tasks are not passed ConfigurationStruct as such need to be initialized with required values
	UpTasks(*parser.ConfigurationStruct) (tasks []Task, err error)

	// devctl up functions
	// Check if any tasks need to be run.
	// Check is also run after Up is finished to ensure function complete successfully
	// Check(*parser.ConfigurationStruct) bool
	// Main runs
	// Up(*parser.ConfigurationStruct)

	// PreInstall(*parser.ConfigurationStruct)
	// Install(*parser.ConfigurationStruct)
	// PostInstall(*parser.ConfigurationStruct)
	PreScript(*parser.ConfigurationStruct)
	Scripts(c *parser.ConfigurationStruct) map[string]utilities.RunCommand
	PostScript(*parser.ConfigurationStruct)
	Down(*parser.ConfigurationStruct)

	// Returns true if plugin applies to the current plugin
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
