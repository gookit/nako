package wex

import (
	"fmt"
	"github.com/gookit/cache"
	"github.com/gookit/ini"
	"github.com/gookit/rux"
	"github.com/gookit/view"
	"github.com/gookit/wex/internal"
	"github.com/syyongx/llog"
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/handler"
	"github.com/syyongx/llog/types"
	"os"
	"time"
)

const (
	EvtBoot   = "app.boot"
	EvtBooted = "app.booted"
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
	Logger *llog.Logger
}

// NewApp new application instance
func NewApp(confFiles ...string) *Application {
	return &Application{
		confFiles: confFiles,

		data: make(map[string]interface{}),

		// services
		Router: rux.New(),
		Config: ini.New(),
	}
}

// Get
func (a *Application) Get() {

}

// Boot application init.
func (a *Application) Boot() {
	var err error

	a.MustFire(EvtBoot, a)

	// load app config
	err = a.Config.LoadExists(a.confFiles...)
	if err != nil {
		panic(err)
	}

	if a.Name == "" {
		a.Name = a.Config.DefString("name", "")
	}

	// views

	a.booted = true
	a.MustFire(EvtBooted, a)
}

func createLogger() {
	logger := llog.NewLogger("wex")

	file := handler.NewFile("/tmp/llog/go.log", 0664, types.WARNING, true)
	buf := handler.NewBuffer(file, 1, types.WARNING, true)
	f := formatter.NewLine("%Datetime% [%LevelName%] [%Channel%] %Message%\n", time.RFC3339)
	file.SetFormatter(f)

	// push handler
	logger.PushHandler(buf)

	// add log
	logger.Warning("xxx")

	// close and write
	buf.Close()
}

// Run the app. addr is optional setting.
// Usage:
//	app.Run()
//	app.Run(":8090")
func (a *Application) Run(addr ...string) {
	if !a.booted {
		a.Boot()
	}

	fmt.Printf("======================== Begin Running(PID: %d) ========================\n", os.Getpid())

	confAddr := a.Config.DefString("listen", "")
	if len(addr) == 0 && confAddr != "" {
		addr = []string{confAddr}
	}

	err := a.Router.Listen(addr...)
	panic(err)
}
