package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rakhazufar/go-project/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func SendJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}

func HashPassword (user *models.User, w http.ResponseWriter) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		message := map[string]string{"message": "Error Creating User"}
		SendJSONResponse(w, http.StatusInternalServerError, message)
		return 	
	}

	user.Password = string(hashPass)
}