package koten

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "show repositories")
}

func TreeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "tree %s!", p.ByName("n"))
}
