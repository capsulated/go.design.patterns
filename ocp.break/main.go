package main

import (
	"github.com/capsulated/workshop.oop.solid.go/ocp.break/server"
	"github.com/capsulated/workshop.oop.solid.go/ocp.break/storage"
	"log"
)

// --== Open Closed Principle BREAK ==--

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	r := server.NewRouter()

	s := storage.NewStorage()

	srv := server.NewServer(r.HttpRouter, s)

	log.Fatal(srv.ListenAndServer())
}
