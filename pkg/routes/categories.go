package routes

//buat CRD untuk Variant
import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/categoriescontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Categories(router *mux.Router) {
	adminApi := router.PathPrefix("/api/v1").Subrouter()

	adminApi.HandleFunc("/create-category", categoriescontrollers.CreateCategory).Methods("POST")
	adminApi.HandleFunc("/get-categories", categoriescontrollers.GetCategories).Methods("GET")
	adminApi.HandleFunc("/delete-category/{id}", categoriescontrollers.DeleteCategoriesById).Methods("DELETE")

	adminApi.Use(middlewares.AdminJWTMiddleware)
}
