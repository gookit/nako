package cmd

import (
	"fmt"
	"github.com/gookit/cliapp"
	"github.com/gookit/cliapp/interact"
	"github.com/gookit/cliapp/show"
	"github.com/gookit/color"
	"os"
	"strings"
)

var createProjectOpts = struct {
	repoUrl string

	//
	appDir string
}{}

// CreateProject command
func CreateProject() *cliapp.Command {
	c := cliapp.NewCommand(
		"create:app",
		"create a new application skeleton project.",
		func(c *cliapp.Command) {
			//
		},
	)

	c.Aliases = []string{"new:app", "new:project"}

	c.StrOpt(&createProjectOpts.repoUrl, "repo-url", "",
		"https://github.com/inhere/go-wex-skeleton",
		"The remote skeleton repo URL address.",
	)
	c.AddArg("name", "the will created project name.", true)
	c.AddArg("dir", "the created project dir path. default is current dir.")
	// dir.Value = "./"

	return c.SetFunc(createProject)
}

// do exec
func createProject(c *cliapp.Command, args []string) int {
	fmt.Println(createProjectOpts)
	fmt.Println(args, c.Arg("dir").String(), c.App().WorkDir())

	name := c.Arg("name").String()
	workDir := c.App().WorkDir()
	targetDir := strings.TrimRight(c.Arg("dir").String(), "/.")
	if targetDir == "" {
		targetDir = workDir
	}

	targetDir += "/" + name

	show.AList("new project info", map[string]string{
		"current dir": workDir,
		"project name": name,
		"project path": targetDir,
		"skeleton url": createProjectOpts.repoUrl,
	}, func(opts *show.ListOption) {
		opts.SepChar = ": "
	})

	if interact.Unconfirmed("ensure create the new project?", true) {
		color.Cyan.Println("Quit create!")
		return 0
	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return c.WithError(err)
	}

	return 0
}
