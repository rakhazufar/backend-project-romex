package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/authcontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Authentication(router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("/login/", authcontrollers.Login).Methods("POST")
	router.HandleFunc("/register/", authcontrollers.Register).Methods("POST")
	api.HandleFunc("/logout/", authcontrollers.Logout).Methods("GET")
	api.Use(middlewares.JWTMiddleware)
}
