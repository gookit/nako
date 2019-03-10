package boot

import (
	"github.com/gookit/event"
	"github.com/gookit/gcli"
	"github.com/gookit/lako"
	"github.com/gookit/lako/cmd/http"
)

// ConsoleBootLoader struct
type ConsoleBootLoader struct {
	Commands []*gcli.Command
}

func (*ConsoleBootLoader) Boot(app *lako.Application) error {
	cliApp := gcli.NewDefaultApp()

	app.MustFire(lako.OnBeforeConsole, event.M{"cliApp": cliApp})

	cliApp.Add(
		http.StartServerCommand(),
		http.RestartServerCommand(),
		http.StopServerCommand(),
	)

	cliApp.Run()

	app.MustFire(lako.OnAfterConsole, event.M{"cliApp": cliApp})

	return nil
}
