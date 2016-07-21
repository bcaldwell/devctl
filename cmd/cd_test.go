package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestCd(t *testing.T) {
	sourceDir, _ := filepath.Abs("../")
	sourceDir = path.Join(sourceDir, "testing_dir", "src")
	pwd, _ := os.Getwd()
	pwd = path.Join(pwd, "..")

	files := getFolderList(sourceDir)

	fileString := fmt.Sprintf("%s", files)

	assert("Failed getFolderList: ", fileString, "["+
		"{"+path.Join(pwd, "/testing_dir")+" src} "+
		"{"+path.Join(pwd, "/testing_dir/src")+" github.com} "+
		"{"+path.Join(pwd, "/testing_dir/src/github.com")+" username} "+
		"{"+path.Join(pwd, "/testing_dir/src/github.com/username")+" golang} "+
		"{"+path.Join(pwd, "/testing_dir/src/github.com/username")+" node}"+
		"]", t)

	match := findMatch("node", files)
	assert("Match failed: ", match, path.Join(pwd, "/testing_dir/src/github.com/username/node"), t)
	match = findMatch("me/node", files)
	assert("Match failed: ", match, path.Join(pwd, "/testing_dir/src/github.com/username/node"), t)
	match = findMatch("go", files)
	assert("Match failed: ", match, path.Join(pwd, "/testing_dir/src/github.com/username/golang"), t)
	match = findMatch("src", files)
	assert("Match failed: ", match, path.Join(pwd, "/testing_dir/src"), t)

}

func assert(pretext, a, b string, t *testing.T) {
	if a != b {
		t.Fatalf("%s expected %s to equal %s", pretext, a, b)
	}
}
