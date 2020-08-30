package boot

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/nico"
	"github.com/gookit/slog"
)

// LogBootLoader struct
type LogBootLoader struct {
}

func (*LogBootLoader) Boot(app *nico.Application) error {
	conf := maputil.MergeStringMap(config.StringMap("log"), map[string]string{
		"name":   "my-log",
		"path":   "/tmp/logs/app.log",
		"level":  "warning",
		"format": "",
		// 0 - disable buffer; >0 - enable buffer
		"bufferSize": "0",
	}, false)

	dump.P(conf)

	app.Logger = slog.New()

	return nil
}
