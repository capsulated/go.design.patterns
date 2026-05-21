package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// --== Single Responsibility Principle SATISFY ==--
// SRP
// structure for each actor ?

type Router struct {
	HttpRouter *httprouter.Router
}

func NewRouter() *Router {
	log.Println("⏳ Initializing router...")

	r := new(Router)
	r.HttpRouter = httprouter.New()

	log.Println("⏳ Initializing handlers...")
	r.HttpRouter.GET("/save", r.saveHandler)

	return r
}

func (*Router) saveHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// DRY BREAK
	//fmt.Fprint(w, "📡 Input request ok")
	//log.Println("📡 Input request ok")

	// Alternative
	reqInitMessage := "📡 Input request initialized"
	fmt.Fprint(w, reqInitMessage)
	log.Println(reqInitMessage)
}

type Storage struct{}

func NewStorage() *Storage {
	return new(Storage)
}

func (*Storage) StoreToFile() {
	log.Println("⏳ Saving to file...")
	// ...
}

func (*Storage) StoreSaveToDb() {
	log.Println("⏳ Inserts to db...")
	// ...
}

type Server struct {
	HttpInstance *http.Server
	Router       *httprouter.Router
	Storage      *Storage
}

func NewServer(r *httprouter.Router, s *Storage) *Server {
	log.Println("⏳ Initializing server...")
	return &Server{
		&http.Server{Addr: ":8080", Handler: r},
		r,
		s,
	}
}

func (s *Server) ListenAndServer() error {
	return s.HttpInstance.ListenAndServe()
}

func (s *Server) Stop() error {
	log.Println("⏳ Stopping server...")
	return s.HttpInstance.Shutdown(nil)
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	router := NewRouter()

	storage := NewStorage()

	server := NewServer(router.HttpRouter, storage)

	log.Fatal(server.ListenAndServer())
}
