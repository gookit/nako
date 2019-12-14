package lako

import (
	"github.com/gookit/cache"
	"github.com/gookit/config/v2"
	"github.com/gookit/rux"
)

var (
	// CtxPool map[string]interface{}
	// storage the global application instance
	defApp = NewApp()
)

// SetGlobal app instance
func SetDefaultApp(app *Application) {
	defApp = app
}

// DefaultApp get application instance
func DefaultApp() *Application {
	return defApp
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
