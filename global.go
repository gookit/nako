package wex

import (
	"github.com/gookit/cache"
	"github.com/gookit/ini"
	"github.com/gookit/rux"
)

// NewGlobalApp create
func NewGlobalApp() *Application {
	_app = NewApp()
	return _app
}

// SetGlobal app instance
func SetGlobal(a *Application)  {
	_app = a
}

// App get application instance
func App() *Application {
	return _app
}

// Router get from global app
func Router() *rux.Router {
	return _app.Router
}

// Router get
func Config() *ini.Ini {
	return _app.Config
}

// Cache get
func Cache() cache.Cache {
	return _app.Cache
}

func Redis() {

}

