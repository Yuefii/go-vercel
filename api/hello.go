package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Go!")
}

func HelloFromID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Hello with ID: %s", id)
}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloWorld).Methods("GET")
	r.HandleFunc("/{id}", HelloFromID).Methods("GET")
	return r
}
