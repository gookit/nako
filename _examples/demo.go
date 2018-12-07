package main

import (
	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
	"github.com/gookit/wex"
)

func main() {
	app := wex.NewApp()

	// add routes
	router := app.Router

	// use middleware
	router.Use(handlers.RequestLogger())

	router.GET("/", func(c *rux.Context) {
		err := c.Text(200, "hello")
		c.Error(err)
	})

	app.Run("localhost:8092")
}

func addRoutes(router *rux.Router)  {

}
