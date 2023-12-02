package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/productcontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Products (router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/products", productcontrollers.CreateProducts).Methods("POST")
	api.HandleFunc("/products", productcontrollers.GetAllProducts).Methods("GET")
	api.HandleFunc("/products/{slug}", productcontrollers.GetProductsBySlug).Methods("GET")
	api.HandleFunc("/products/{slug}", productcontrollers.EditProductsBySlug).Methods("PUT")
	api.HandleFunc("/products/{slug}", productcontrollers.DeleteProductsBySlug).Methods("DELETE")
	api.Use(middlewares.AdminJWTMiddleware)
}


