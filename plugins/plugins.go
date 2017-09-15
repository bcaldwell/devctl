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
	// Note: tasks are not passed ProjectConfig as such need to be initialized with required values
	UpTasks(*parser.ProjectConfig) (tasks []Task, err error)

	// devctl up functions
	// Check if any tasks need to be run.
	// Check is also run after Up is finished to ensure function complete successfully
	// Check(*parser.ProjectConfig) bool
	// Main runs
	// Up(*parser.ProjectConfig)

	// PreInstall(*parser.ProjectConfig)
	// Install(*parser.ProjectConfig)
	// PostInstall(*parser.ProjectConfig)
	PreScript(*parser.ProjectConfig)
	Scripts(c *parser.ProjectConfig) map[string]utilities.RunCommand
	PostScript(*parser.ProjectConfig)
	Down(*parser.ProjectConfig)

	// Returns true if plugin applies to the current plugin
	IsProjectType(*parser.ProjectConfig) bool
}

// List is a list of all available plugins
var List = []Plugin{
	&Docker{},
	// &Golang{},
}

// Used calls IsProjectType on each element on the PluginList array and returns a filtered version
func Used(c *parser.ProjectConfig) []Plugin {
	pluginsUsed := []Plugin{}

	for _, plugin := range List {
		if plugin.IsProjectType(c) {
			pluginsUsed = append(pluginsUsed, plugin)
		}
	}
	return pluginsUsed
}
