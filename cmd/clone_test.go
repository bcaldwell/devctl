package cmd

import (
	"encoding/json"
	"testing"

	"github.com/bcaldwell/devctl/parser"
)

type User struct {
	name string
}

type testStruct struct {
	args     []string
	config   *cloneConfig
	expected *cloneConfig
}

func TestCloneGithub(t *testing.T) {

	parser.DevctlConfig = &parser.Config{
		GithubUser: "github_user",
		SourceDir:  "/",
	}

	testStructs := []testStruct{
		{
			[]string{"test"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "test",
				User:      "github_user",
				Host:      "github.com",
				Url:       "git@github.com:github_user/test",
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
				Url:       "git@github.com:username/test",
				SourceDir: "/src/github.com/username",
			},
		},
	}

	for _, i := range testStructs {
		i.config.github()
		i.config.parseArgs(i.args)
		i.config.setSourceDir()
		i.config.setUrl()

		checkStruct(t, &i)
	}

}

func TestCloneGitlabWithoutURLConfigured(t *testing.T) {

	parser.DevctlConfig = &parser.Config{
		GithubUser: "github_user",
		SourceDir:  "/",
		GitlabUser: "gitlab_user",
	}

	testStructs := []testStruct{
		{
			[]string{"test"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "test",
				User:      "gitlab_user",
				Host:      "gitlab.com",
				Url:       "git@gitlab.com:gitlab_user/test",
				SourceDir: "/src/gitlab.com/gitlab_user",
			},
		},
		{
			[]string{"username/test"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "test",
				User:      "username",
				Host:      "gitlab.com",
				Url:       "git@gitlab.com:username/test",
				SourceDir: "/src/gitlab.com/username",
			},
		},
	}

	for _, i := range testStructs {
		i.config.github()
		i.config.gitlab()
		i.config.parseArgs(i.args)
		i.config.setSourceDir()
		i.config.setUrl()

		checkStruct(t, &i)
	}

}

func TestCloneGitlabConfigured(t *testing.T) {

	parser.DevctlConfig = &parser.Config{
		GithubUser: "github_user",
		SourceDir:  "/",
		GitlabURL:  "gitlab.somwhere.com",
		GitlabUser: "gitlab_user",
	}

	testStructs := []testStruct{
		{
			[]string{"test"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "test",
				User:      "gitlab_user",
				Host:      "gitlab.somwhere.com",
				Url:       "git@gitlab.somwhere.com:gitlab_user/test",
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
				Url:       "git@gitlab.somwhere.com:username/test",
				SourceDir: "/src/gitlab.somwhere.com/username",
			},
		},
	}

	for _, i := range testStructs {
		i.config.github()
		i.config.gitlab()
		i.config.parseArgs(i.args)
		i.config.setSourceDir()
		i.config.setUrl()

		checkStruct(t, &i)
	}

}

func TestCloneFullUrlWithoutGit(t *testing.T) {

	parser.DevctlConfig = &parser.Config{
		GithubUser: "github_user",
		SourceDir:  "/",
		GitlabURL:  "gitlab.somwhere.com",
		GitlabUser: "gitlab_user",
	}

	testStructs := []testStruct{
		{
			[]string{"https://github.com/user1/project1"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "project1",
				User:      "user1",
				Host:      "github.com",
				Url:       "https://github.com/user1/project1",
				SourceDir: "/src/github.com/user1",
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

func TestCloneFullUrl(t *testing.T) {

	parser.DevctlConfig = &parser.Config{
		GithubUser: "github_user",
		SourceDir:  "/",
		GitlabURL:  "gitlab.somwhere.com",
		GitlabUser: "gitlab_user",
	}

	testStructs := []testStruct{
		{
			[]string{"https://github.com/user1/project1.git"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "project1",
				User:      "user1",
				Host:      "github.com",
				Url:       "https://github.com/user1/project1.git",
				SourceDir: "/src/github.com/user1",
			},
		},
		{
			[]string{"git@github.com:user1/project1.git"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "project1",
				User:      "user1",
				Host:      "github.com",
				Url:       "git@github.com:user1/project1.git",
				SourceDir: "/src/github.com/user1",
			},
		},
		{
			[]string{"git@gitlab.somewhere.com:user2/project3.git"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "project3",
				User:      "user2",
				Host:      "gitlab.somewhere.com",
				Url:       "git@gitlab.somewhere.com:user2/project3.git",
				SourceDir: "/src/gitlab.somewhere.com/user2",
			},
		},
		{
			[]string{"http://gitlab.somewhere.com/user3/project.git"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "project",
				User:      "user3",
				Host:      "gitlab.somewhere.com",
				Url:       "http://gitlab.somewhere.com/user3/project.git",
				SourceDir: "/src/gitlab.somewhere.com/user3",
			},
		},
		{
			[]string{"git@github.com:user1/project1.git"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "project1",
				User:      "user1",
				Host:      "github.com",
				Url:       "git@github.com:user1/project1.git",
				SourceDir: "/src/github.com/user1",
			},
		},
		{
			[]string{"git@github.com:user1/project1.git"},
			new(cloneConfig),
			&cloneConfig{
				Repo:      "project1",
				User:      "user1",
				Host:      "github.com",
				Url:       "git@github.com:user1/project1.git",
				SourceDir: "/src/github.com/user1",
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
// func TestCloneGitlabPartlyConfigured(t *testing.T) {

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

// func TestCloneGitlabUnconfigured(t *testing.T) {

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
