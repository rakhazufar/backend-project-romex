package routes

//buat CRD untuk Variant
import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/statuscontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Status(router *mux.Router) {
	adminApi := router.PathPrefix("/api/v1").Subrouter()

	adminApi.HandleFunc("/get-status", statuscontrollers.GetStatus).Methods("GET")

	adminApi.Use(middlewares.AdminJWTMiddleware)
}
