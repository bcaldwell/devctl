package postCommand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter(t *testing.T) {

	type tester struct {
		dir      []string
		commands []string
		output   string
	}

	testData := []tester{
		{
			[]string{
				"/test",
			},
			[]string{},
			"cd /test",
		},
		{
			[]string{},
			[]string{
				"nvm use 4",
			},
			"nvm use 4",
		},
		{
			[]string{
				"/test",
			},
			[]string{
				"nvm use 4",
			},
			"cd /test\nnvm use 4",
		},
		{
			[]string{
				"/test",
			},
			[]string{
				"nvm use 4",
				"nvm use 6",
			},
			"cd /test\nnvm use 4\nnvm use 6",
		},
		{
			[]string{
				"/test",
				"/demo",
			},
			[]string{
				"python test.py",
			},
			"cd /test\ncd /demo\npython test.py",
		},
	}

	for _, test := range testData {
		for _, cd := range test.dir {
			ChangeDir(cd)
		}
		for _, cmd := range test.commands {
			RunCommand(cmd)
		}
		assert.Equal(t, writeString(), test.output, "")
		reset()
	}
}

func TestWriterAlternate(t *testing.T) {
	ChangeDir("/test")
	RunCommand("nvm use 4")
	ChangeDir("/demo")
	RunCommand("npm install")

	assert.Equal(t, writeString(), "cd /test\nnvm use 4\ncd /demo\nnpm install")
	reset()
}

func reset() {
	p.commands = p.commands[:0]
}
