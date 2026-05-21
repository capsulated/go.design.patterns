package main

import (
	"github.com/capsulated/workshop.oop.solid.go/isp.break/storage"
	"log"
)

// --== Interface Segrigation Principle BREAK ==--
// ISP
// open for extension, closed for modification

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	_ = storage.NewStorage()
}
