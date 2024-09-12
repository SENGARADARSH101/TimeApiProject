package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", getcurrentTime)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
