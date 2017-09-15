package plugins

import (
	"os"
	"path"

	"strings"

	"github.com/benjamincaldwell/devctl/parser"
	"github.com/benjamincaldwell/devctl/postCommand"
	"github.com/benjamincaldwell/devctl/shell"
	"github.com/benjamincaldwell/devctl/utilities"
	"github.com/benjamincaldwell/go-printer"
)

func init() {
	virtualEnvLocation = path.Join(os.Getenv("HOME"), "/.devctl", utilities.LocationHash(), "venv")
}

var virtualEnvLocation string

const virtualEnvLink = ".devctl/venv"

// Dependency on pip

type Python struct {
	pipBinary    string
	pythonBinary string
	path         string
}

func (p *Python) Setup() {
	// printer.Info("Setting up python")
	p.dectectSystemBinary()
	p.pipInstallIfNotInstalled("virtualenv")
}

func (p *Python) PreInstall(c *parser.ProjectConfig) {
	p.dectectSystemBinary()
	if err := p.pipInstallIfNotInstalled("virtualenv"); err != nil {
		printer.Fail("Couldn't find or install virtualenv package")
	} else {
		if _, err := os.Stat(virtualEnvLocation); os.IsNotExist(err) {
			printer.Info("Creating virtualenv")

			pythonBinaryVersioned := "python"
			if c.Python.Version == "3" {
				pythonBinaryVersioned = "python3"
			} else if c.Python.Version == "2" {
				pythonBinaryVersioned = "python2"
			} else {
				printer.Warning("Invalid python version. Using system python")
			}

			err := shell.Command("virtualenv", "-p", pythonBinaryVersioned, virtualEnvLocation).SetPath(p.path).Run()
			utilities.ErrorCheck(err, "created virtualenv")
		} else {
			printer.Success("virtualenv already created")
		}

		if _, err := os.Stat(virtualEnvLink); os.IsNotExist(err) {
			printer.Info("creating link to virtualenv")
			if dir, err := os.Getwd(); err == nil {
				err = os.Symlink(virtualEnvLocation, path.Join(dir, virtualEnvLink))
				utilities.HandleError(err)
			} else {
				utilities.HandleError(err)
			}
		}
	}

}

func (n *Python) Install(c *parser.ProjectConfig) {
	if _, err := os.Stat("requirements.txt"); os.IsNotExist(err) {
		printer.Success("not requirements to install")
	} else {
		err := shell.Command("sh", "-c", "source "+virtualEnvLocation+"/bin/activate; pip install -r requirements.txt").PrintOutput()
		utilities.ErrorCheck(err, "requirements install")
	}
}

func (n Python) PostInstall(c *parser.ProjectConfig) {
	postCommand.RunCommand("source " + virtualEnvLocation + "/bin/activate")
}

func (n Python) PreScript(c *parser.ProjectConfig) {
}

func (n Python) Scripts(c *parser.ProjectConfig) map[string]utilities.RunCommand {
	// return scripts struct array
	scripts := make(map[string]utilities.RunCommand)

	scripts["start"] = utilities.RunCommand{
		Command: "python __main__.py",
	}

	return scripts
}

func (n Python) PostScript(c *parser.ProjectConfig) {
}

func (n Python) Down(c *parser.ProjectConfig) {
	postCommand.RunCommand("deactivate")
}

func (n Python) IsProjectType(c *parser.ProjectConfig) bool {
	if c.Python.Version != "" {
		return true
	}
	return false
}

// dectectSystemBinary sets pythonBinary and pipBinary to system installed versions. Also sets python path without virtualenv set
func (p *Python) dectectSystemBinary() {
	if venv := os.Getenv("VIRTUAL_ENV"); venv != "" {
		printer.Warning("Running in virtualenv. Trying to detect system python installation")

		p.path = utilities.RemoveFromPath(venv + "/bin")
	} else {
		p.path = os.Getenv("PATH")
	}
	if output, err := shell.Command("which", "python").SetPath(p.path).Output(); err == nil {
		p.pythonBinary = strings.TrimSpace(string(output))
	} else {
		printer.Error("Couldn't find system python")
	}
	if output, err := shell.Command("which", "pip").SetPath(p.path).Output(); err == nil {
		p.pipBinary = strings.TrimSpace(string(output))
	} else {
		printer.Error("Couldn't find system pip")
	}
}

func (p Python) pipCheckIfInstalled(packageName string) bool {
	err := shell.Command(p.pipBinary, "show", packageName).Run()
	return (err == nil)
}

func (p Python) pipInstallIfNotInstalled(packageName string) error {
	if p.pipCheckIfInstalled("virtualenv") {
		printer.Success("%s already installed", packageName)
	} else {
		printer.Info("Installing %s", packageName)
		printer.InfoLineTop()
		err := shell.Command("sudo", p.pipBinary, "install", packageName).PrintOutput()
		printer.InfoLineBottom()
		utilities.ErrorCheck(err, "install virtualenv")
		return err
	}
	return nil
}
