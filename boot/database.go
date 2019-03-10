package boot

import "github.com/gookit/lako"

type DBBootLoader struct {
}

func (*DBBootLoader) Boot(app *lako.Application) error {

	return nil
}
