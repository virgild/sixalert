package server

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"goji.io/pat"

	"github.com/virgild/sixalert/agent"
	"github.com/virgild/sixalert/alert"
	"github.com/virgild/sixalert/utils"
	"github.com/virgild/sixalert/websocket"

	"github.com/gorilla/handlers"
	"github.com/tylerb/graceful"
	"github.com/urfave/cli"
	"goji.io"
)

type AppServer struct {
	*graceful.Server
	Mux             *goji.Mux
	SocketServer    *websocket.Server
	AlertsCh        <-chan *alert.TTCAlert
	DevelopmentMode bool
}

// Returns a new AppServer
func NewAppServer(devMode bool) *AppServer {
	mux := goji.NewMux()

	srv := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":3000",
			Handler: handlers.LoggingHandler(os.Stdout, mux),
		},
	}

	socketServer := websocket.NewServer()

	s := &AppServer{
		Server:          srv,
		SocketServer:    socketServer,
		Mux:             mux,
		AlertsCh:        make(<-chan *alert.TTCAlert),
		DevelopmentMode: devMode,
	}

	s.setupHandlers()

	return s
}

// Set up the app server handlers
func (s *AppServer) setupHandlers() {
	if s.DevelopmentMode {
		dir, err := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0]), "..", "web/dist"))
		if err != nil {
			log.Fatal(err)
		}

		s.Mux.HandleFunc(pat.Get("/"), func(w http.ResponseWriter, r *http.Request) {
			file := filepath.Join(dir, "index.html")
			http.ServeFile(w, r, file)
		})

		s.Mux.HandleFunc(pat.Get("/bundle.js"), func(w http.ResponseWriter, r *http.Request) {
			file := filepath.Join(dir, "bundle.js")
			http.ServeFile(w, r, file)
		})
	} else {
		html, _ := Asset("web/dist/index.html")
		bundle, _ := Asset("web/dist/bundle.js")
		s.Mux.HandleFunc(pat.Get("/"), func(w http.ResponseWriter, r *http.Request) {
			w.Write(html)
		})

		s.Mux.HandleFunc(pat.Get("/bundle.js"), func(w http.ResponseWriter, r *http.Request) {
			w.Write(bundle)
		})
	}
}

func (s *AppServer) Run() {
	if s.DevelopmentMode {
		go runWebpackWatch()
	}
	go s.SocketServer.Listen(s.Mux)
	go func() {
		for {
			select {
			case alert := <-s.AlertsCh:
				log.Printf("%v", alert)
			}
		}
	}()
	s.ListenAndServe()
}

func runWebpackWatch() {
	output := log.New(os.Stdout, "WEBPACK: ", log.LstdFlags)
	logOutput := utils.LoggerIO(output)
	cmd := exec.Command("npm", "run", "watch")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	cmd.Dir = filepath.Join(dir, "..", "web")

	cmd.Stdout = logOutput
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func ServerCommand() cli.Command {
	cmd := cli.Command{
		Name:  "server",
		Usage: "Starts the web service",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "dev",
				Usage: "Development mode (loads assets from file, starts webpack watcher)",
			},
		},
		Action: func(c *cli.Context) error {
			var err error
			alertsCh := make(<-chan *alert.TTCAlert)
			agent := agent.NewAgent(alertsCh)
			go agent.Start()
			app := NewAppServer(c.Bool("dev"))
			app.AlertsCh = alertsCh
			app.Run()

			return err
		},
	}
	return cmd
}
