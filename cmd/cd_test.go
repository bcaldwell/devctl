package cmd

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCd(t *testing.T) {
	sourceDir, _ := filepath.Abs("../")
	sourceDir = path.Join(sourceDir, "testing_dir", "src")
	pwd, _ := os.Getwd()
	pwd = path.Join(pwd, "..")

	folders := getFolderList(sourceDir)

	expectedFolders := []folder{
		{
			path.Join(pwd, "/testing_dir"),
			"src",
		},
		{
			path.Join(pwd, "/testing_dir/src"),
			"github.com",
		},
		{
			path.Join(pwd, "/testing_dir/src/github.com"),
			"username",
		},
		{
			path.Join(pwd, "/testing_dir/src/github.com/username"),
			"node",
		},
		{
			path.Join(pwd, "/testing_dir/src/github.com/username"),
			"golang",
		},
	}

	assert.Equal(t, len(folders), len(expectedFolders), "Number of folders returned")

	for _, folder := range expectedFolders {
		assert.Contains(t, folders, folder, "Value of folders returned")
	}

	match := findMatch("node", folders)
	assert.Equal(t, match, path.Join(pwd, "/testing_dir/src/github.com/username/node"), "Match failed: ")
	match = findMatch("username/node", folders)
	assert.Equal(t, match, path.Join(pwd, "/testing_dir/src/github.com/username/node"), "Match failed: ")
	match = findMatch("go", folders)
	assert.Equal(t, match, path.Join(pwd, "/testing_dir/src/github.com/username/golang"), "Match failed: ")
	match = findMatch("src", folders)
	assert.Equal(t, match, path.Join(pwd, "/testing_dir/src"), "Match failed: ")
}
