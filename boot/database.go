package boot

import (
	"github.com/gookit/nico"
)

type DBBootLoader struct {
}

func (*DBBootLoader) Boot(app *nico.Application) error {

	return nil
}
