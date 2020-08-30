package boot

import (
	"github.com/gookit/nako"
)

type DBBootLoader struct {
}

func (*DBBootLoader) Boot(app *nako.Application) error {

	return nil
}
