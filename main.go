package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rakhazufar/go-project/pkg/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	r := mux.NewRouter()
	routes.Authentication(r)
	routes.Address(r)
	routes.Admin(r)
	routes.Products(r)

	log.Fatal(http.ListenAndServe(":9010", r))
}