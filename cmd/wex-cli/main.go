package main

import (
	"github.com/gookit/gcli"
	"github.com/gookit/gcli/builtin"
	"github.com/gookit/wex/cmd"
	"runtime"
)

// run:
// go run ./cmd/wex-cli
// go build ./cmd/wex-cli && ./wex-cli
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := gcli.NewApp(func(app *gcli.App) {
		app.Version = "1.0.6"
		app.Description = "this is wex cli application"
		app.Hooks[gcli.EvtInit] = func(a *gcli.App, data interface{}) {
			// do something...
			// fmt.Println("init app")
		}
		// app.SetVerbose(cliapp.VerbDebug)
		// app.DefaultCommand("example")
		app.Logo.Text = `   ________    _______
  / ____/ /   /  _/   |  ____  ____
 / /   / /    / // /| | / __ \/ __ \
/ /___/ /____/ // ___ |/ /_/ / /_/ /
\____/_____/___/_/  |_/ .___/ .___/
                     /_/   /_/`
	})

	app.Add(
		cmd.CreateProjectCommand(),
		builtin.GenAutoCompleteScript(),
	)

	app.Run()
}
