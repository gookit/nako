package main

import (
	"github.com/gookit/event"
	"github.com/gookit/nico"
	"github.com/gookit/nico/boot"
	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
)

// go build _examples/demo.go && ./demo
func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	app := nico.DefaultApp()

	// add routes
	router := app.Router

	// use middleware
	router.Use(handlers.RequestLogger())

	router.GET("/", func(c *rux.Context) {
		c.Text(200, "hello")
	})

	app.On(nico.OnAfterBoot, event.ListenerFunc(func(e event.Event) error {
		return nil
	}))

	app.BootLoaders = []nico.BootLoader{
		boot.EnvBootLoader("./", ".env"),
		boot.ConfigBootLoader("./_examples/config.yml"),
		&boot.LogBootLoader{},
		&boot.ConsoleBootLoader{},
	}

	app.Run()
}
