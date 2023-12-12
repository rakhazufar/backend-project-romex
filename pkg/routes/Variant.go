package routes

//buat CRD untuk Variant
import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/variantcontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Variant(router *mux.Router) {
	adminApi := router.PathPrefix("/api/v1").Subrouter()

	adminApi.HandleFunc("/variant/{id}", variantcontrollers.DeleteVariantById).Methods("DELETE")

	adminApi.Use(middlewares.AdminJWTMiddleware)
}
