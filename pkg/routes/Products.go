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
	api.HandleFunc("/products/{productsSlug}", productcontrollers.GetProductsByName).Methods("GET")
	api.HandleFunc("/products/{productsId}", productcontrollers.EditProductsById).Methods("PUT")
	api.HandleFunc("/products/{productsSlug}", productcontrollers.DeleteProductsById).Methods("GET")
	api.Use(middlewares.AdminJWTMiddleware)
}


