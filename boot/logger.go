package boot

import (
	"github.com/gookit/config"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/lako"
	"github.com/syyongx/llog"
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/handler"
	"github.com/syyongx/llog/types"
	"time"
)

// LogBootLoader struct
type LogBootLoader struct {

}

func (*LogBootLoader) Boot(app *lako.Application) error {
	conf = maputil.MergeStringMap(config.StringMap("log"), map[string]string{
		"name":   "my-log",
		"path":   "/tmp/logs/app.log",
		"level":  "warning",
		"format": "",
		// 0 - disable buffer; >0 - enable buffer
		"bufferSize": "0",
	}, false)

	logger := llog.NewLogger("lako")

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

	return nil
}
