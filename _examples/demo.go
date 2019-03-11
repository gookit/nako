package main

import (
	"github.com/gookit/event"
	"github.com/gookit/lako"
	"github.com/gookit/lako/boot"
	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
)

// go build _examples/demo.go && demo
func main() {
	app := lako.DefaultApp()

	// add routes
	router := app.Router

	// use middleware
	router.Use(handlers.RequestLogger())

	router.GET("/", func(c *rux.Context) {
		c.Text(200, "hello")
	})

	app.On(lako.OnAfterBoot, event.ListenerFunc(func(e event.Event) error {
		return nil
	}))

	app.BootLoaders = []lako.BootLoader{
		boot.EnvBootLoader("./", ".env"),
		boot.ConfigBootLoader("./config/app.ini"),
		&boot.LogBootLoader{},
		&boot.ConsoleBootLoader{},
	}

	app.Run()
}
