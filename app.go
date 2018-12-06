package wex

import (
	"github.com/gookit/cache"
	"github.com/gookit/ini"
	"github.com/gookit/rux"
	"github.com/gookit/view"
)

var (
	CtxPool map[string]interface{}
	// storage the application instance
	_app *application
)

// application instance
type application struct {
	Name   string
	View   *view.Renderer
	Cache cache.Cache
	Config *ini.Ini
	Router *rux.Router
}

// NewApp new application instance
func NewApp(confFiles ...string) *application {
	return &application{}
}

// App get application instance
func App() *application {
	return _app
}

// Boot application
func Boot(confFiles ...string) {
	var err error

	// load app config
	Config, err = ini.LoadExists(confFiles...)
	if err != nil {
		panic(err)
	}

	// cache
}

// Router get
func Router() *rux.Router {
	return _app.Router
}

// Router get
func Config() *ini.Ini {
	return _app.Config
}

func Cache() cache.Cache {
	return nil
}

func Redis() {

}
