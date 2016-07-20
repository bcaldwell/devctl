package cmd

import (
	"encoding/json"
	"testing"

	"github.com/spf13/viper"
)

type User struct {
	name string
}

type testStruct struct {
	args     []string
	config   *cloneConfig
	expected *cloneConfig
}

func TestCdGithub(t *testing.T) {

	viper.SetDefault("github_user", "github_user")
	viper.SetDefault("source_dir", "/")

	testStructs := []testStruct{
		{
			[]string{"test"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "test",
				User:      "github_user",
				Host:      "github.com",
				Url:       "github.com",
				SourceDir: "/src/github.com/github_user",
			},
		},
		{
			[]string{"username/test"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "test",
				User:      "username",
				Host:      "github.com",
				Url:       "github.com",
				SourceDir: "/src/github.com/username",
			},
		},
	}

	for _, i := range testStructs {
		i.config.github()
		i.config.parseArgs(i.args)
		i.config.setSourceDir()

		checkStruct(t, &i)
	}

}

func TestCdGitlabConfigured(t *testing.T) {

	viper.SetDefault("github_user", "github_user")
	viper.SetDefault("source_dir", "/")
	viper.SetDefault("gitlab_url", "gitlab.somwhere.com")
	viper.SetDefault("gitlab_user", "gitlab_user")

	testStructs := []testStruct{
		{
			[]string{"test"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "test",
				User:      "gitlab_user",
				Host:      "gitlab.somwhere.com",
				Url:       "gitlab.somwhere.com",
				SourceDir: "/src/gitlab.somwhere.com/gitlab_user",
			},
		},
		{
			[]string{"username/test"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "test",
				User:      "username",
				Host:      "gitlab.somwhere.com",
				Url:       "gitlab.somwhere.com",
				SourceDir: "/src/gitlab.somwhere.com/username",
			},
		},
	}

	for _, i := range testStructs {
		i.config.github()
		i.config.gitlab()
		i.config.parseArgs(i.args)
		i.config.setSourceDir()

		checkStruct(t, &i)
	}

}

// ya this doesnt work....
// func TestCdGitlabPartlyConfigured(t *testing.T) {

// 	viper.SetDefault("github_user", "github_user")
// 	viper.SetDefault("source_dir", "/")
// 	viper.SetDefault("gitlab_url", "gitlab.somwhere.com")
// 	viper.SetDefault("gitlab_user", "gitlab_user")

// 	testStructs := []testStruct{
// 		{
// 			[]string{"test"},
// 			new(cloneConfig),
// 			&cloneConfig{
// 				Repo:      "test",
// 				User:      "gitlab_user",
// 				Host:      "gitlab.somwhere.com",
// 				Url:       "gitlab.somwhere.com",
// 				SourceDir: "/src/gitlab.somwhere.com/gitlab_user",
// 			},
// 		},
// 	}

// 	for _, i := range testStructs {
// 		i.config.github()
// 		i.config.gitlab()
// 		i.config.parseArgs(i.args)
// 		i.config.setSourceDir()

// 		checkStruct(t, &i)
// 	}

// }

// func TestCdGitlabUnconfigured(t *testing.T) {

// 	viper.SetDefault("github_user", "github_user")
// 	viper.SetDefault("source_dir", "/")
// 	// viper.SetDefault("gitlab_url", "")
// 	// viper.SetDefault("gitlab_user", "")

// 	testStructs := []testStruct{
// 		{
// 			[]string{"test"},
// 			new(cloneConfig),
// 			&cloneConfig{
// 				Repo:      "test",
// 				User:      "github_user",
// 				Host:      "github.com",
// 				Url:       "github.com",
// 				SourceDir: "",
// 			},
// 		},
// 		{
// 			[]string{"username/test"},
// 			new(cloneConfig),
// 			&cloneConfig{
// 				Repo:      "test",
// 				User:      "username",
// 				Host:      "github.com",
// 				Url:       "github.com",
// 				SourceDir: "",
// 			},
// 		},
// 	}

// 	for _, i := range testStructs {
// 		i.config.github()
// 		i.config.gitlab()
// 		i.config.parseArgs(i.args)
// 		i.config.setSourceDir()

// 		checkStruct(t, &i)
// 	}

// }

func checkStruct(t *testing.T, s *testStruct) {
	configBytes, _ := json.Marshal(s.config)
	expectedBytes, _ := json.Marshal(s.expected)

	config := string(configBytes)
	expected := string(expectedBytes)

	if config != expected {
		t.Fatalf("Failed with args: %s \n Expected %s to equal %s", s.args, config, expected)
	}
}
