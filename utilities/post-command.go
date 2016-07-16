package utilities

import "os"

// PostCommand struct
type PostCommand struct {
	dir string
}

// ChangeDir sets the directory to cd to after command exits
func (p *PostCommand) ChangeDir(dirPath string) {
	p.dir = dirPath
}

// Write writes to the file descriptor
func (p *PostCommand) Write() {
	fd := os.NewFile(8, "fd")
	fd.Write([]byte(p.dir))

	fd.Write([]byte("\n"))
	fd.Close()

}
