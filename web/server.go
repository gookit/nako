package web

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gookit/event"
	"github.com/gookit/lako"
)

// HTTPServer an HTTP web server
type HTTPServer struct {
	srv *http.Server

	pidFile  string
	address  []string
	realAddr string

	processID int
}

// NewHTTPServer create new HTTPServer.
// Usage:
// 	srv := NewHTTPServer("127.0.0.1")
// 	srv := NewHTTPServer("127.0.0.1:8090")
// 	srv := NewHTTPServer("127.0.0.1", "8090")
func NewHTTPServer(address ...string) *HTTPServer {
	return &HTTPServer{
		processID: os.Getpid(),

		address:  address,
		realAddr: resolveAddress(address),
	}
}

/*************************************************************
 * Start HTTP server
 *************************************************************/

// Start server, begin handle HTTP request
func (s *HTTPServer) Start() error {
	app := lako.App()

	s.srv = &http.Server{
		Addr: s.realAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(500)
				}
			}()

			evtData := event.M{"w": w, "r": r}

			// Fire before route
			app.MustFire(OnBeforeRoute, evtData)

			// Route and dispatch request
			app.Router.ServeHTTP(w, r)

			// Fire after route
			app.MustFire(OnAfterRoute, evtData)
		}),
	}

	app.MustFire(OnServerStart, event.M{"addr": s.srv.Addr})

	// listen signal
	s.handleSignal(s.srv)

	err := s.srv.ListenAndServe()

	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

/*************************************************************
 * Getter/Setter methods
 *************************************************************/

// RealAddr get resolved read addr
func (s *HTTPServer) RealAddr() string {
	return s.realAddr
}

// ProcessID return
func (s *HTTPServer) ProcessID() int {
	return s.processID
}

// PidFile get pid file path
func (s *HTTPServer) PidFile() string {
	return s.pidFile
}

// SetPidFile set pid file path
func (s *HTTPServer) SetPidFile(pidFile string) {
	s.pidFile = pidFile
}

// handleSignal handles system signal for graceful shutdown.
func (s *HTTPServer) handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		fmt.Printf("Got signal [%s], exiting server now", s)
		if err := server.Close(); err != nil{
			fmt.Printf("Server close failed: %s", err.Error())
		}

		lako.App().MustFire(OnServerClose, event.M{"sig": s})
		// service.DisconnectDB()

		fmt.Println("Server exited")
		os.Exit(0)
	}()
}

func removePidFile(pidFile string) error {
	return os.Remove(pidFile)
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
