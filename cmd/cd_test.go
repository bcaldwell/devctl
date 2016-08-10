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
			"devctl",
		},
		{
			path.Join(pwd, "/testing_dir/src/github.com/devctl"),
			"mean-example",
		},
		{
			path.Join(pwd, "/testing_dir/src/github.com/devctl"),
			"docker-compose-example",
		},
		{
			path.Join(pwd, "/testing_dir/src/github.com/devctl"),
			"golang-example",
		},
	}

	assert.Equal(t, len(expectedFolders), len(folders), "Number of folders returned")

	for _, folder := range expectedFolders {
		assert.Contains(t, folders, folder, "Value of folders returned")
	}

	match := findMatch("mean", folders)
	assert.Equal(t, match, path.Join(pwd, "/testing_dir/src/github.com/devctl/mean-example"), "Match failed: ")
	match = findMatch("dev/mea", folders)
	assert.Equal(t, match, path.Join(pwd, "/testing_dir/src/github.com/devctl/mean-example"), "Match failed: ")
	match = findMatch("compose", folders)
	assert.Equal(t, match, path.Join(pwd, "/testing_dir/src/github.com/devctl/docker-compose-example"), "Match failed: ")
	match = findMatch("src", folders)
	assert.Equal(t, match, path.Join(pwd, "/testing_dir/src"), "Match failed: ")
}
