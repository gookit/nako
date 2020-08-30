package nico

import (
	"github.com/gookit/cache"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/gookit/event"
	"github.com/gookit/rux"
	"github.com/gookit/slog"
	"github.com/gookit/view"
)

// Application instance
type Application struct {
	*event.Manager

	Name string
	data map[string]interface{}

	BootLoaders []BootLoader

	// components
	View   *view.Renderer
	Cache  cache.Cache
	Config *config.Config
	Router *rux.Router
	Logger *slog.Logger
}

// NewApp new application instance
func NewApp() *Application {
	app := &Application{
		data: make(map[string]interface{}),

		// services
		Router: rux.New(),
		Config: config.New("lako"),
		// events
		Manager: event.NewManager("gweb"),
	}

	// add yaml support
	app.Config.AddDriver(yaml.Driver)

	return app
}

// Run the application
// Usage:
// 	app.Run()
func (a *Application) Run() {
	a.MustFire(OnBeforeBoot, event.M{"app": a})

	a.bootstrap()

	a.MustFire(OnAfterBoot, event.M{"app": a})
}

// Bootstrap application init.
func (a *Application) bootstrap() {
	for _, loader := range a.BootLoaders {
		if err := loader.Boot(a); err != nil {
			panic(err)
		}
	}

	if a.Name == "" {
		a.Name = a.Config.String("name", "")
	}
}

// Set value to app.data
func (a *Application) Set(name string, val interface{}) {
	a.data[name] = val
}

// Get value from app.data
func (a *Application) Get(name string) interface{} {
	if val, ok := a.data[name]; ok {
		return val
	}
	return nil
}
