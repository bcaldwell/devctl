package plugins

// import (
// 	"testing"

// 	"github.com/benjamincaldwell/devctl/parser"
// 	"github.com/stretchr/testify/assert"
// )

// func TestPluginsUsed(t *testing.T) {
// 	type testCases struct {
// 		config   string
// 		expected []Plugin
// 	}

// 	tests := []testCases{
// 		{
// 			`{
// 				"Node": {
// 					"Version": "4"
// 				}
// 			}`,
// 			[]Plugin{&Node{}},
// 		},
// 		{
// 			`{
// 			}`,
// 			[]Plugin{},
// 		},
// 	}

// 	for _, test := range tests {
// 		config := parser.ProjectConfigStruct{}
// 		config.ParseJson(test.config)
// 		assert.Equal(t, Used(&config), test.expected, "")
// 	}
// }
