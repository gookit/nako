package main

import (
	"github.com/gookit/event/simpleevent"
	"github.com/gookit/lako"
	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
)

func main() {

	app := lako.NewApp()

	// add routes
	router := app.Router

	// use middleware
	router.Use(handlers.RequestLogger())

	router.GET("/", func(c *rux.Context) {
		c.Text(200, "hello")
	})

	app.On("http.run", func(e *simpleevent.EventData) error {
		// httpSrv.Run()
		return nil
	})

	app.Run("localhost:8092")
}

func addRoutes(router *rux.Router) {

}
