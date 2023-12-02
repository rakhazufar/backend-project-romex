package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/imagecontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Image (router *mux.Router) {
	adminApi := router.PathPrefix("/api/v1").Subrouter()
	adminApi.HandleFunc("/image", imagecontrollers.UploadImage).Methods("POST")
	adminApi.Use(middlewares.AdminJWTMiddleware)
}