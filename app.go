package lako

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gookit/cache"
	"github.com/gookit/config"
	"github.com/gookit/event"
	"github.com/gookit/lako/boot"
	"github.com/gookit/rux"
	"github.com/gookit/view"
	"github.com/syyongx/llog"
)

var (
	// CtxPool map[string]interface{}
	// storage the global application instance
	defApp *Application
)

// Application instance
type Application struct {
	*event.Manager

	Name string
	data map[string]interface{}

	bootLoaders []BootLoader

	booted bool

	confFiles []string
	// components
	View   *view.Renderer
	Cache  cache.Cache
	Config *config.Config
	Router *rux.Router
	Logger *llog.Logger
}

// NewApp new application instance
func NewApp(confFiles ...string) *Application {
	app := &Application{
		confFiles: confFiles,

		data: make(map[string]interface{}),

		// services
		Router: rux.New(),
		Config: config.New("lako"),
		// events
		Manager: event.NewManager("lako"),
	}

	return app
}

// Get
func (a *Application) defaultLoaders() []BootLoader {
	return []BootLoader{
		&boot.LogBootLoader{},
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

// Bootstrap application init.
func (a *Application) bootstrap() {
	var err error

	a.MustFire(BeforeBoot, event.M{"app": a})

	// load app config
	err = a.Config.LoadExists(a.confFiles...)
	if err != nil {
		panic(err)
	}

	if a.Name == "" {
		a.Name = a.Config.String("name", "")
	}

	// views
	a.booted = true
	a.MustFire(AfterBoot, event.M{"app": a})
}

// Run the app. addr is optional setting.
// Usage:
// 	app.Run()
// 	app.Run(":8090")
func (a *Application) Run(addr ...string) {

	a.bootstrap()

	fmt.Printf("======================== Begin Running(PID: %d) ========================\n", os.Getpid())

	confAddr := a.Config.String("listen", "")
	if len(addr) == 0 && confAddr != "" {
		addr = []string{confAddr}
	}

	a.Router.Listen(addr...)
}

func (a *Application) BootLoaders() []BootLoader {
	return a.bootLoaders
}

func (a *Application) SetBootLoaders(bootLoaders []BootLoader) {
	a.bootLoaders = bootLoaders
}

/*************************************************************
 * handle HTTP request
 *************************************************************/

// ServeHTTP handle HTTP request
func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(500)
		}
	}()

	a.MustFire(AfterBoot, event.M{"w": w, "r": r})

	a.Router.ServeHTTP(w, r)

	a.MustFire(AfterBoot, event.M{"w": w, "r": r})
}
