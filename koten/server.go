package koten

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func registerHandlers(r *httprouter.Router) {
	r.GET("/", ListHandler)
	r.GET("/repos/:n", TreeHandler)
}

func Run() {
	r := httprouter.New()
	registerHandlers(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
