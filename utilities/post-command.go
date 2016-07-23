package utilities

import "os"

// PostCommand struct
type PostCommand struct {
	commands []string
}

// ChangeDir sets the directory to cd to after command exits
func (p *PostCommand) ChangeDir(dirPath string) {
	p.commands = append(p.commands, "cd "+dirPath)
}

func (p *PostCommand) RunCommand(command string) {
	p.commands = append(p.commands, command)
}

// Write writes to the file descriptor
func (p *PostCommand) Write() {
	fd := os.NewFile(8, "fd")

	for _, command := range p.commands {
		fd.Write([]byte(command))
		fd.Write([]byte("\n"))
	}

	fd.Close()

}
