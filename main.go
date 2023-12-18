package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/routes"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()

	routes.Authentication(r)
	routes.Address(r)
	routes.Admin(r)
	routes.Products(r)
	routes.Image(r)
	routes.Variant(r)
	routes.Categories(r)
	routes.Token(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           86400,
	})

	// Terapkan CORS ke router
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":9010", handler))
}
