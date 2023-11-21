package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/addresscontrollers"
)

func Address (router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/address", addresscontrollers.PostAddress).Methods("POST")
	api.HandleFunc("/address", addresscontrollers.GetAddress).Methods("GET")
}