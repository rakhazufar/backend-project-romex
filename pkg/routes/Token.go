package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/tokencontrollers"
)

func Token(router *mux.Router) {
	router.HandleFunc("/verify-token", tokencontrollers.VerifyToken).Methods("POST")

}
