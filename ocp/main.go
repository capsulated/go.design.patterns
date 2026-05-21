package main

import (
	"github.com/capsulated/workshop.oop.solid.go/ocp/server"
	"github.com/capsulated/workshop.oop.solid.go/ocp/storage"
	"log"
)

// --== Open / Closed Principle SATISFY ==--
// OCP
// open for extension, closed for modification

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	r := server.NewRouter()

	s := storage.NewStorage()

	srv := server.NewServer(r.HttpRouter, s)

	log.Fatal(srv.ListenAndServer())
}
