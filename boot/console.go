package boot

import (
	"github.com/gookit/gcli"
	"github.com/gookit/lako"
)

// ConsoleBootLoader struct
type ConsoleBootLoader struct {
	Commands []*gcli.Command
}

func (*ConsoleBootLoader) Boot(app *lako.Application) error {

	return nil
}
