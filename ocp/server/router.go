package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Router struct {
	HttpRouter *httprouter.Router
}

func NewRouter() *Router {
	log.Println("⏳ Initializing router...")

	r := new(Router)
	r.HttpRouter = httprouter.New()

	log.Println("⏳ Initializing handlers...")
	r.HttpRouter.GET("/save", r.HandleSave)

	return r
}

func (*Router) HandleSave(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// DRY BREAK
	fmt.Fprint(w, "📡 Input request initialized")
	log.Println("📡 Input request initialized")

	// Alternative
	// reqInitMessage := "📡 Input request initialized"
	// fmt.Fprint(w, reqInitMessage)
	// log.Println(reqInitMessage)
}
