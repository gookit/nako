package cmd

import "github.com/gookit/gcli"

// StopServerCommand stop server
func StopServerCommand() *gcli.Command {
	return &gcli.Command{
		Name: "serve:stop",
		UseFor: "stop the running http server",
		Aliases: []string{"http:stop"},
		Func: func(c *gcli.Command, args []string) error{
			return nil
		},
	}
}
