package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/admincontrollers"
)

func Admin (router *mux.Router) {
	router.HandleFunc("/login/", admincontrollers.AdminLogin).Methods("POST")
	router.HandleFunc("/register/", admincontrollers.AdminRegister).Methods("POST")
}


