package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePageHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
