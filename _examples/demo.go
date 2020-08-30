package main

import (
	"github.com/gookit/event"
	"github.com/gookit/nako"
	"github.com/gookit/nako/boot"
	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
)

// go build _examples/demo.go && ./demo
func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	app := nako.DefaultApp()

	// add routes
	router := app.Router

	// use middleware
	router.Use(handlers.RequestLogger())

	router.GET("/", func(c *rux.Context) {
		c.Text(200, "hello")
	})
	router.GET("/routes", func(c *rux.Context) {
		c.Text(200, router.String())
	})

	app.On(nako.OnAfterBoot, event.ListenerFunc(func(e event.Event) error {
		return nil
	}))

	app.BootLoaders = []nako.BootLoader{
		boot.EnvBootLoader("./", ".env"),
		boot.ConfigBootLoader("./_examples/config.yml"),
		&boot.LogBootLoader{},
		&boot.ConsoleBootLoader{},
	}

	app.Run()
}
