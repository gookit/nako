package boot

import (
	"fmt"

	"github.com/gookit/config/v2/dotnev"
	"github.com/gookit/event"
	"github.com/gookit/nico"
)

// EnvBootLoader for load .env file
func EnvBootLoader(dir string, envFiles ...string) nico.BootLoader {
	return nico.BootFunc(func(app *nico.Application) error {
		return dotnev.LoadExists(dir, envFiles...)
	})
}

// EnvBootLoader for load config files
func ConfigBootLoader(confFiles ...string) nico.BootLoader {
	return nico.BootFunc(func(app *nico.Application) error {
		app.MustFire(nico.OnBeforeConfig, event.M{
			"files":  confFiles,
			"config": app.Config,
		})

		fmt.Println("load config files:", confFiles)

		// load from files
		err := app.Config.LoadExists(confFiles...)

		// load from flags
		if err == nil {
			err = app.Config.LoadFlags([]string{"debug"})
		}

		app.MustFire(nico.OnAfterConfig, event.M{"config": app.Config})

		return err
	})
}
