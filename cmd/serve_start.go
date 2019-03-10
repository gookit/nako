package cmd

import (
	"fmt"
	"github.com/gookit/gcli"
	"github.com/gookit/lako"
	"github.com/gookit/lako/web"
	"os"
)

var httpServeOpts = struct {
	addr     string

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

	confAddr := lako.Config().String("listen", "0.0.0.0:8080")

	c.StrOpt(&httpServeOpts.addr, "addr", "s", confAddr,
		"The HTTP server listen address",
	)

	c.Func = func(c *gcli.Command, args []string) error {
		fmt.Printf("======================== Begin Running(PID: %d) ========================\n", os.Getpid())

		web.HTTPServer{}.Run(httpServeOpts.addr)

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
	c.StrOpt(&httpServeOpts.appDir, "repo-url", "",
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
