package utils

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

func SendJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}

func HashPassword(userpass string) (*string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(userpass), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashedAdminPasswword := string(hashPass)

	return &hashedAdminPasswword, nil
}

func HashAdminPassword(adminpass string) (*string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(adminpass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashedAdminPasswword := string(hashPass)

	return &hashedAdminPasswword, nil
}

func Slugify(title string) string {
	u := uuid.New()
	shortUUID := u.String()[:5]
	slugTitle := slug.Make(title + "-" + shortUUID)
	
	return slugTitle
}