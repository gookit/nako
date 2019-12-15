package main

import (
	"io/ioutil"
	"runtime"

	"github.com/gookit/gcli/v2"
	"github.com/gookit/gcli/v2/builtin"
	"github.com/gookit/nico/cmd"
)

// run:
// go run ./cmd/lako
// go build ./cmd/lako && ./lako
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := gcli.NewApp(func(app *gcli.App) {
		app.Version = "1.0.6"
		app.Description = "this is lako cli application"
		app.On(gcli.EvtInit, func(data ...interface{}) {
			// do something...
			// fmt.Println("init app")
		})

		// app.SetVerbose(gcli.VerbDebug)
		// app.DefaultCommand("example")
		textBts, _ := ioutil.ReadFile("resource/fontlogo/ansi-shadow.txt")
		app.Logo.Text = string(textBts)
	})

	app.Add(
		cmd.CreateProjectCommand(),
		builtin.GenAutoCompleteScript(),
	)

	app.Run()
}
