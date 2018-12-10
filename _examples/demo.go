package main

import (
	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
	"github.com/gookit/wex"
	"net/http"
)

func main() {
	app := wex.NewApp()

	// add routes
	router := app.Router

	// use middleware
	router.Use(handlers.RequestLogger())

	router.GET("/", func(c *rux.Context) {
		c.Text(200, "hello")
	})

	app.BeforeRoute = func(w http.ResponseWriter, r *http.Request) {

	}

	app.Run("localhost:8092")
}

func addRoutes(router *rux.Router) {

}
