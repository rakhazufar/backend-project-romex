package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/authcontrollers"
)

func Authentication (router *mux.Router) {
	router.HandleFunc("/login/", authcontrollers.Login).Methods("POST")
	router.HandleFunc("/register/", authcontrollers.Register).Methods("POST")
	router.HandleFunc("/logout/", authcontrollers.Logout).Methods("GET")
}
