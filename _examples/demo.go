package main

import (
	"github.com/gookit/event"
	"github.com/gookit/lako"
	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
)

func main() {
	app := lako.NewDefaultApp()

	// add routes
	router := app.Router

	// use middleware
	router.Use(handlers.RequestLogger())

	router.GET("/", func(c *rux.Context) {
		c.Text(200, "hello")
	})

	app.On(lako.AfterBoot, event.ListenerFunc(func(e event.Event) error {
		return nil
	}))

	app.Run("localhost:8092")
}

func addRoutes(router *rux.Router) {

}
