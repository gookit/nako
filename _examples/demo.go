package main

import (
	"github.com/gookit/rux"
	"github.com/gookit/wex"
)

func main() {
	app := wex.NewApp()

	addRoutes(app.Router)

	app.Run()
}

func addRoutes(router *rux.Router)  {
	router.GET("/", func(c *rux.Context) {
		err := c.Text(200, "hello")
		c.Error(err)
	})
}
