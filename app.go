package wex

import (
	"github.com/gookit/cache"
	"github.com/gookit/ini"
	"github.com/gookit/rux"
	"github.com/gookit/view"
	"github.com/gookit/wex/internal"
)

var (
	// CtxPool map[string]interface{}
	// storage the global application instance
	_app *Application
)

// Application instance
type Application struct {
	internal.SimpleEvent

	Name string
	data map[string]interface{}

	booted bool

	confFiles []string

	View   *view.Renderer
	Cache  cache.Cache
	Config *ini.Ini
	Router *rux.Router
}

// NewApp new application instance
func NewApp(confFiles ...string) *Application {
	return &Application{
		confFiles: confFiles,

		data: make(map[string]interface{}),
	}
}

// Get
func (a *Application) Get() {

}

// Boot application init.
func (a *Application) Boot() {
	var err error

	// load app config
	a.Config, err = ini.LoadExists(a.confFiles...)
	if err != nil {
		panic(err)
	}

	if a.Name == "" {
		a.Name = a.Config.DefString("name", "")
	}

	a.booted = true
}

// Run app
func (a *Application) Run() {
	if !a.booted {
		a.Boot()
	}

	err := a.Router.Listen(":80")
	panic(err)
}

