package http

import "github.com/gookit/gcli"

var httpServeOpts = struct {
	repoUrl  string
	forceNew bool

	//
	appDir string
}{}

func StartServerCommand() *gcli.Command {
	c := gcli.NewCommand(
		"serve:start",
		"create a new application skeleton project.",
		func(c *gcli.Command) {
			//
		},
	)

	c.Aliases = []string{"serve", "http:start"}

	// zip bag https://github.com/inhere/go-web-skeleton/archive/master.zip
	c.StrOpt(&httpServeOpts.repoUrl, "repo-url", "",
		"https://github.com/inhere/go-web-skeleton",
		"The remote skeleton repo URL address.",
	)
	c.BoolOpt(&httpServeOpts.forceNew, "force-new", "f", false,
		"force re-download repo archive package",
	)

	c.Func = func(c *gcli.Command, args []string) error {
		return nil
	}

	return c
}

func RestartServerCommand() *gcli.Command {
	c := gcli.NewCommand(
		"serve:restart",
		"create a new application skeleton project.",
		func(c *gcli.Command) {
			//
		},
	)

	c.Aliases = []string{"http:restart"}

	// zip bag https://github.com/inhere/go-web-skeleton/archive/master.zip
	c.StrOpt(&httpServeOpts.repoUrl, "repo-url", "",
		"https://github.com/inhere/go-web-skeleton",
		"The remote skeleton repo URL address.",
	)
	c.BoolOpt(&httpServeOpts.forceNew, "force-new", "f", false,
		"force re-download repo archive package",
	)

	c.Func = func(c *gcli.Command, args []string) error {
		return nil
	}

	return c
}

