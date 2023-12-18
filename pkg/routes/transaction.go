package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/controllers/transactioncontrollers"
	"github.com/rakhazufar/go-project/pkg/middlewares"
)

func Transaction(router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("/create-transaction", transactioncontrollers.CreateTransaction).Methods("POST")
	api.HandleFunc("/get-transaction", transactioncontrollers.GetTransactions).Methods("GET")
	api.HandleFunc("/get-transaction/{id}", transactioncontrollers.GetTransactionById).Methods("GET")
	api.HandleFunc("/update-transaction", transactioncontrollers.UpdateTransaction).Methods("PATCH")
	api.Use(middlewares.AdminJWTMiddleware)
}
