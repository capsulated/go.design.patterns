package main

import "log"

// --== Single Responsibility BREAK ==-- ?

type Server struct{}

func (*Server) Init() {
	log.Println("⏳ Initializing server...")
	// Realisation
}

func (*Server) ListenAndServer() {
	log.Println("⏳ Starting serve...")
	// Realisation
}

func (*Server) Stop() {
	log.Println("⏳ Stopping server...")
	// Realisation
}

// Comment: God Object
func (*Server) SetHandlers() {
	log.Println("⏳ Initializing handlers...")
	// ...
}

func (s *Server) saveHandler() {
	log.Println("📡 Input request initiated")
	if true {
		s.StoreSaveToDb()
	} else {
		s.StoreToFile()
	}
	// ...
}

func (*Server) StoreToFile() {
	log.Println("⏳ Saving file...")
	// ...
}

func (*Server) StoreSaveToDb() {
	log.Println("⏳ Inserts file...")
	// ...
}

// Comment: God Object END

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	s := new(Server)

	s.Init()
	s.SetHandlers()

	go s.ListenAndServer()
}
