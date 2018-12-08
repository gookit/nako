package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/gookit/gcli"
	"github.com/gookit/gcli/helper"
	"github.com/gookit/gcli/interact"
	"github.com/gookit/gcli/show"
	"github.com/gookit/goutil/fsUtil"
	"strings"
)

var createProjectOpts = struct {
	repoUrl  string
	forceNew bool

	//
	appDir string
}{}

// CreateProjectCommand command
// run: wex-cli new:app demo ./test
func CreateProjectCommand() *gcli.Command {
	c := gcli.NewCommand(
		"create:app",
		"create a new application skeleton project.",
		func(c *gcli.Command) {
			//
		},
	)

	c.Aliases = []string{"new:app", "new:project"}

	// zip bag https://github.com/inhere/go-wex-skeleton/archive/master.zip
	c.StrOpt(&createProjectOpts.repoUrl, "repo-url", "",
		"https://github.com/inhere/go-wex-skeleton",
		"The remote skeleton repo URL address.",
	)
	c.BoolOpt(&createProjectOpts.forceNew, "force-new", "f", false,
		"force re-download repo archive package",
	)
	c.AddArg("name", "the will created project name.", true)
	c.AddArg("dir", "the created project dir path. default is current dir.")
	// dir.Value = "./"

	return c.SetFunc(createProject)
}

// do exec
func createProject(c *gcli.Command, args []string) (err error) {
	fmt.Println(createProjectOpts)
	fmt.Println(args, c.Arg("dir").String(), c.App().WorkDir())

	name := c.Arg("name").String()
	workDir := c.App().WorkDir()
	targetDir := strings.TrimRight(c.Arg("dir").String(), "/.")
	if targetDir == "" {
		targetDir = workDir
	}

	projectDir := targetDir + "/" + name

	show.AList("new project info", map[string]string{
		"current dir":  workDir,
		"project name": name,
		"project path": projectDir,
		"skeleton url": createProjectOpts.repoUrl,
	}, func(opts *show.ListOption) {
		opts.SepChar = ": "
	})

	if interact.Unconfirmed("ensure create the new project?", true) {
		color.Cyan.Println("Quit create!")
		return
	}

	err = downloadZIPArchive(
		createProjectOpts.repoUrl+"/archive/master.zip",
		"./",
		"skeleton-archive.zip",
	)
	if err != nil {
		return
	}

	err = fsUtil.Unzip("./skeleton-archive.zip", targetDir)
	return
}

// https://github.com/inhere/go-wex-skeleton/archive/master.zip
func downloadZIPArchive(url, saveDir, filename string) (err error) {
	return helper.Download(url, saveDir, filename)
}
