package nico

const (
	ModeDev  mode = "dev"
	ModeTest mode = "test"
	ModeProd mode = "prod"
)

var (
	Debug = false
	Mode  = ModeDev
)

type mode string

// BootLoader for app start boot
type BootLoader interface {
	// Boot do something before application run
	Boot(app *Application) error
}

// BootFunc for application
type BootFunc func(app *Application) error

// Boot do something
func (fn BootFunc) Boot(app *Application) error {
	return fn(app)
}

// SetMode set run mode
func SetMode(name mode) {
	Mode = name
}
