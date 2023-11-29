package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/admincontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Admin (router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	adminApi := router.PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("/admin/login", admincontrollers.AdminLogin).Methods("POST")
	api.HandleFunc("/admin/create", admincontrollers.AdminRegister).Methods("POST")
	adminApi.HandleFunc("/admin/logout", admincontrollers.Logout).Methods("GET")
	api.Use(middlewares.AdministratorJWTMiddleware)
	adminApi.Use(middlewares.AdminJWTMiddleware)
}


