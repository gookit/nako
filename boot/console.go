package boot

import (
	"github.com/gookit/event"
	"github.com/gookit/gcli/v2"
	"github.com/gookit/nako"
	"github.com/gookit/nako/cmd"
)

// ConsoleBootLoader struct
type ConsoleBootLoader struct {
	Commands []*gcli.Command
}

func (*ConsoleBootLoader) Boot(app *nako.Application) error {
	cliApp := gcli.NewApp()

	app.MustFire(nako.OnBeforeConsole, event.M{"cliApp": cliApp})

	cliApp.Add(
		cmd.StartServerCommand(),
		cmd.RestartServerCommand(),
		cmd.StopServerCommand(),
	)

	cliApp.Run()

	app.MustFire(nako.OnAfterConsole, event.M{"cliApp": cliApp})

	return nil
}
