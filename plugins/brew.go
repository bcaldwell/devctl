package plugins

import (
	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/benjamincaldwell/go-printer"
)

// Brew plugin struct
type Brew struct {
	dependencies []string
}

// Setup for brew installs brew if it is not installed
func (b *Brew) Setup() {
	if utilities.CheckIfInstalled("brew") {
		printer.Success("brew already installed")
	} else {
		printer.Info("Installing brew using offical install scripts")
		printer.InfoLineTop()
		err := shell.Command("sh", "-c", "/usr/bin/ruby -e \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"").PrintOutput()
		printer.InfoLineBottom()
		utilities.HandleError(err)
	}
}

func (b *Brew) PreInstall(c *parser.ProjectConfig) {
	b.dependencies = append(c.Dependencies.Install, c.Dependencies.Brew.Install...)
}

func (b *Brew) Install(c *parser.ProjectConfig) {
	for _, dependency := range b.dependencies {
		if !checkIfBrewInstalled(dependency) {
			brewInstall(dependency)
		} else {
			printer.VerboseSuccess("%s already installed", dependency)
		}
	}
}

func (b Brew) PostInstall(c *parser.ProjectConfig) {
}

func (b Brew) PreScript(c *parser.ProjectConfig) {
}

func (b Brew) Scripts(c *parser.ProjectConfig) map[string]utilities.RunCommand {
	scripts := make(map[string]utilities.RunCommand)
	return scripts
}

func (b Brew) PostScript(c *parser.ProjectConfig) {
}

func (b Brew) Down(c *parser.ProjectConfig) {
}

func (n Brew) IsProjectType(c *parser.ProjectConfig) bool {
	if len(c.Dependencies.Install) > 0 || len(c.Dependencies.Brew.Install) > 0 {
		return true
	}
	return false
}

// checkIfBrewInstalled checks if a formulae is installed using brew. Returns true if it is installed
func checkIfBrewInstalled(formulae string) bool {
	err := shell.Command("brew", "ls", formulae).Run()
	return (err == nil)
}

func brewInstall(formulae string) error {
	printer.Info("Installing %s", formulae)
	printer.InfoLineTop()
	err := shell.Command("brew", "isntall", formulae).PrintOutput()
	printer.InfoLineBottom()
	return err
}
