package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/addresscontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Address(router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/address", addresscontrollers.CreateAddress).Methods("POST")
	api.HandleFunc("/address", addresscontrollers.UpdateAddress).Methods("PUT")
	api.HandleFunc("/address", addresscontrollers.GetAddress).Methods("GET")
	api.Use(middlewares.JWTMiddleware)
}
