package routes

//buat CRD untuk Variant
// import (
// 	"github.com/gorilla/mux"
// 	"github.com/rakhazufar/go-project/pkg/middlewares"
// )

// func Variant(router *mux.Router) {
// 	adminApi := router.PathPrefix("/api/v1").Subrouter()

// 	// adminApi.HandleFunc("/categories", categoriescontrollers.CreateVariant).Methods("POST")
// 	// adminApi.HandleFunc("/categories", categoriescontrollers.GetAllVariant).Methods("GET")
// 	// adminApi.HandleFunc("/categories/{id}", categoriescontrollers.DeleteCategoriesById).Methods("DELETE")

// 	adminApi.Use(middlewares.AdminJWTMiddleware)
// }
