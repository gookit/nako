package lako

import (
	"github.com/gookit/cache"
	"github.com/gookit/config"
	"github.com/gookit/rux"
)

// NewGlobalApp create
func NewDefaultApp(confFiles ...string) *Application {
	defApp = NewApp(confFiles...)
	return defApp
}

// SetGlobal app instance
func SetDefaultApp(app *Application) {
	defApp = app
}

// App get application instance
func App() *Application {
	return defApp
}

// Router get from global app
func Router() *rux.Router {
	return defApp.Router
}

// Router get
func Config() *config.Config {
	return defApp.Config
}

// Cache get
func Cache() cache.Cache {
	return defApp.Cache
}

func Redis() {

}
