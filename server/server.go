package server

import (
	"log"
	"net/http"

	"github.com/aodin/volta/config"

	v1 "github.com/aodin/groupthink/server/api/v1"
)

const VERSION = "0.0.1"

type Server struct {
	Config config.Config
}

// ListenAndServe starts the server and serves forever
func (server *Server) ListenAndServe() error {
	log.Printf("server: serving on address %s\n", server.Config.Address())
	return http.ListenAndServe(server.Config.Address(), nil)
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	return
}

func New(conf config.Config) *Server {
	http.Handle("/", http.RedirectHandler(v1.Prefix, 302))
	http.HandleFunc("/favicon.ico", Favicon)
	http.HandleFunc(v1.Prefix, v1.Query)
	return &Server{Config: conf}
}
