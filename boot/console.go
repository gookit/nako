package boot

import (
	"github.com/gookit/event"
	"github.com/gookit/gcli/v2"
	"github.com/gookit/nico"
	"github.com/gookit/nico/cmd"
)

// ConsoleBootLoader struct
type ConsoleBootLoader struct {
	Commands []*gcli.Command
}

func (*ConsoleBootLoader) Boot(app *nico.Application) error {
	cliApp := gcli.NewApp()

	app.MustFire(nico.OnBeforeConsole, event.M{"cliApp": cliApp})

	cliApp.Add(
		cmd.StartServerCommand(),
		cmd.RestartServerCommand(),
		cmd.StopServerCommand(),
	)

	cliApp.Run()

	app.MustFire(nico.OnAfterConsole, event.M{"cliApp": cliApp})

	return nil
}
