package isp

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Storager interface {
	Store()
	Delete()
}

type FileStorer interface {
	FileStore()
}

type DbStorer interface {
	DbStore()
}

type UniversalStorage interface {
	Storager
	FileStorer
	DbStorer
}

type FileStorage struct{}

type DbStorage struct{}

func NewFileStorage() *FileStorage {
	return new(Storage)
}

func (*FileStorage) Store() {
	log.Println("⏳ Saving...")
	// ...
}

func (*FileStorage) Delete() {
	log.Println("⏳ Delete...")
	// ...
}

func (*FileStorage) FileStore() {
	log.Println("⏳ Delete...")
	// ...
}

func NewDbStorage() *DbStorage {
	return new(Storage)
}

func (*DbStorage) Store() {
	log.Println("⏳ Saving...")
	// ...
}

func (*DbStorage) Delete() {
	log.Println("⏳ Delete...")
	// ...
}

func (*DbStorage) DbStore() {
	log.Println("⏳ Delete...")
	// ...
}

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

func (*Router) saveHandler(ctx) {
	reqInitMessage := "📡 Input request initialized"
	fmt.Fprint(w, reqInitMessage)
	log.Println(reqInitMessage)
}

type Server struct {
	HttpInstance *http.Server
	Router       *Router
	Storage      *Storage
}

func NewServer(r *httprouter.Router, s Storager) *Server {
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

func main() {

	r := NewRouter()

	c := NewCache()

	fs := NewFileStorage(c)

	dbs := NewDbStorage(c)

	s := NewServer(r.HttpRouter, fs, c)

	go s.ListenAndServer()
}
