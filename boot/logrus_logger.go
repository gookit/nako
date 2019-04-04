package boot

import (
	"github.com/gookit/lako"
	"github.com/sirupsen/logrus"
)

// RFC3339NanoFixed is time.RFC3339Nano with nanoseconds padded using zeros to
// ensure the formatted time isalways the same number of characters.
// copy form docker codes.
const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"

// LogrusBootLoader struct
type LogrusBootLoader struct {
	//
}

// Boot Logrus component
func (lb *LogrusBootLoader) Boot(app *lako.Application) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: RFC3339NanoFixed,
		DisableColors:   false,
		FullTimestamp:   true,
	})

	// logrus.JSONFormatter{}
	return nil
}
