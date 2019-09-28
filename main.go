package main

import (
	"net/http"

    "github.com/gorilla/mux"
	"github.com/hisamura333/todo/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Index).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}