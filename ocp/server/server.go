package server

import (
	"github.com/capsulated/workshop.oop.solid.go/ocp/storage"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Server struct {
	HttpInstance *http.Server
	Router       *httprouter.Router
	Storage      *storage.Storage
}

func NewServer(r *httprouter.Router, s *storage.Storage) *Server {
	log.Println("⏳ Initializing server...")
	return &Server{
		&http.Server{Addr: ":8080", Handler: r},
		r,
		s,
	}
}

func (s *Server) ListenAndServer() error {
	log.Println("⭐️ Start server ⭐️")
	return s.HttpInstance.ListenAndServe()
}

func (s *Server) Stop() error {
	log.Println("⏳ Stopping server...")
	return s.HttpInstance.Shutdown(nil)
}
