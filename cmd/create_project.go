package cmd

import (
	"fmt"
	"github.com/gookit/cliapp"
	"github.com/gookit/cliapp/interact"
	"github.com/gookit/cliapp/progress"
	"github.com/gookit/cliapp/show"
	"github.com/gookit/cliapp/utils"
	"github.com/gookit/color"
	"github.com/gookit/goutil/fs"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var createProjectOpts = struct {
	repoUrl  string
	forceNew bool

	//
	appDir string
}{}

// CreateProjectCommand command
// run: wex-cli new:app demo ./test
func CreateProjectCommand() *cliapp.Command {
	c := cliapp.NewCommand(
		"create:app",
		"create a new application skeleton project.",
		func(c *cliapp.Command) {
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
func createProject(c *cliapp.Command, args []string) int {
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
		return 0
	}

	err := downloadZIPArchive(
		createProjectOpts.repoUrl+"/archive/master.zip",
		"./",
		"skeleton-archive.zip",
	)
	if err != nil {
		return c.WithError(err)
	}

	err = fs.Unzip("./skeleton-archive.zip", targetDir)
	if err != nil {
		return c.WithError(err)
	}

	return 0
}

// https://github.com/inhere/go-wex-skeleton/archive/master.zip
func downloadZIPArchive(url, saveDir, filename string) (err error) {
	return utils.Download(url, saveDir, filename)
}
