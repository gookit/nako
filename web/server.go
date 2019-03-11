package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gookit/event"
	"github.com/gookit/lako"
)

// Logger
var logger *log.Logger

// HTTPServer an HTTP web server
type HTTPServer struct {
	srv *http.Server
}

// NewHTTPServer
func NewHTTPServer() *HTTPServer {
	logger = log.New(os.Stdout, "", log.LstdFlags)
	return &HTTPServer{}
}

/*************************************************************
 * handle HTTP request
 *************************************************************/

// Run handle HTTP request
func (s *HTTPServer) Run(addr ...string) {
	app := lako.App()

	s.srv = &http.Server{
		Addr: resolveAddress(addr),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(500)
				}
			}()

			app.MustFire(OnBeforeRoute, event.M{"w": w, "r": r})

			app.Router.ServeHTTP(w, r)

			app.MustFire(OnAfterRoute, event.M{"w": w, "r": r})
		}),
	}

	app.MustFire(OnServerStart, event.M{"addr": s.srv.Addr})

	// listen signal
	s.handleSignal(s.srv)

	err := s.srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// handleSignal handles system signal for graceful shutdown.
func (s *HTTPServer) handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logger.Printf("Got signal [%s], exiting server now", s)
		if err := server.Close(); nil != err {
			logger.Printf("Server close failed: %s", err.Error())
		}

		lako.App().MustFire(OnServerClose, event.M{"sig": s})
		// service.DisconnectDB()

		logger.Println("Server exited")
		os.Exit(0)
	}()
}

func resolveAddress(addr []string) (fullAddr string) {
	ip := "0.0.0.0"
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); len(port) > 0 {
			fmt.Printf("Environment variable PORT=\"%s\"", port)
			return ip + ":" + port
		}
		fmt.Printf("Environment variable PORT is undefined. Using port :8080 by default")
		return ip + ":8080"
	case 1:
		var port string
		if strings.IndexByte(addr[0], ':') != -1 {
			ss := strings.SplitN(addr[0], ":", 2)
			if ss[0] != "" {
				return addr[0]
			}
			port = ss[1]
		} else {
			port = addr[0]
		}

		return ip + ":" + port
	default:
		panic("too much parameters")
	}
}
