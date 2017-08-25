package systemDetector

import (
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
	"strings"
)

type System struct {
	Name       string
	PrettyName string
	ID         string
	Version    string
}

func DetectSystem() (sys System, err error) {
	if runtime.GOOS != "linux" {
		sys.ID = runtime.GOOS
		return sys, err
	}
	osReleaseFile := "/etc/os-release"
	if _, err := os.Stat(osReleaseFile); err == nil {
		fileContents, err := ioutil.ReadFile(osReleaseFile)
		if err != nil {
			return sys, err
		}
		parseReleaseFile(string(fileContents), &sys)
	}

	archFile := "/etc/arch-release"
	if _, err := os.Stat(archFile); err == nil {
		sys.ID = "arch"
		sys.Name = "Arch Linux"
	}

	return sys, err
}

func parseReleaseFile(fileContents string, sys *System) {
	r := regexp.MustCompile(`\bNAME="(.*?)"`)
	sys.Name = parseline(r.FindString(fileContents))

	r = regexp.MustCompile(`\bPRETTY_NAME="(.*?)"`)
	sys.PrettyName = parseline(r.FindString(fileContents))

	r = regexp.MustCompile(`\bID="?(.*?)"|\bID=(.*?)\n`)
	sys.ID = parseline(r.FindString(fileContents))

	r = regexp.MustCompile(`\bVERSION="(.*?)"`)
	sys.Version = parseline(r.FindString(fileContents))
}

func parseline(line string) string {
	if line == "" {
		return line
	}
	value := strings.Split(line, "=")[1]
	return strings.TrimSpace(strings.Replace(value, "\"", "", -1))
}
