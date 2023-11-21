package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Authentication(r)
	routes.Address(r)

	log.Fatal(http.ListenAndServe(":9010", r))
}