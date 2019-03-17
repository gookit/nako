package cmd

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/gookit/gcli"
	"github.com/gookit/lako"
	"github.com/gookit/lako/web"
)

var httpServeOpts = struct {
	addr string

	forceNew bool

	//
	appDir string
}{}

func StartServerCommand() *gcli.Command {
	c := &gcli.Command{
		Name:    "serve:start",
		UseFor:  "start the new http server",
		Aliases: []string{"http:start"},
	}

	confAddr := lako.Config().String("listen", "0.0.0.0:8080")
	c.StrOpt(&httpServeOpts.addr, "addr", "s", confAddr,
		"The HTTP server listen address",
	)

	c.Func = func(c *gcli.Command, args []string) error {
		return startServer()
	}

	return c
}

// StopServerCommand stop server
func StopServerCommand() *gcli.Command {
	return &gcli.Command{
		Name:    "serve:stop",
		UseFor:  "stop the running http server",
		Aliases: []string{"http:stop"},
		Func: func(c *gcli.Command, args []string) error {
			err := stopServer()
			if err == nil {
				 color.Success.Println("Server stopped")
			}
			return err
		},
	}
}

func RestartServerCommand() *gcli.Command {
	c := &gcli.Command{
		Name: "serve:restart",
		UseFor: "restart the running http server",
		Aliases: []string{"http:restart"},
	}

	confAddr := lako.Config().String("listen", "0.0.0.0:8080")
	c.StrOpt(&httpServeOpts.addr, "addr", "s", confAddr,
		"The HTTP server listen address",
	)

	c.Func = func(c *gcli.Command, args []string) error {
		srv := createServer()
		if srv.IsRunning() { // Stop old
			err := srv.Stop(3)
			if err != nil {
				return err
			}
		}

		return startServer()
	}
	return c
}

func createServer() *web.HTTPServer  {
	srv := web.NewHTTPServer(httpServeOpts.addr)
	srv.SetPidFile(lako.Config().String("pidFile"))

	return srv
}

func startServer() error {
	srv := createServer()
	if srv.IsRunning() {
		return fmt.Errorf("cannot start, server is already running(PID: %d)", srv.ProcessID())
	}

	addr := srv.RealAddr()

	fmt.Printf("======================== Begin Running(PID: %d) ========================\n", srv.ProcessID())
	color.Printf("Serve listen on %s Go to http://%s\n", addr, addr)

	return srv.Start()
}

func stopServer() error {
	srv := createServer()
	if srv.IsRunning() {
		return srv.Stop(3)
	}

	pid := srv.ProcessID()
	return fmt.Errorf("cannot stop, the server is not running(PID: %d)", pid)
}