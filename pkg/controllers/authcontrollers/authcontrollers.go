package authcontrollers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/rakhazufar/go-project/pkg/config"
	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	//membaca json dari r.body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		message := map[string]string{"message": "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}

	defer r.Body.Close()

	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

	user, err := models.GetUserByUsername(userInput.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			message := map[string]string{"message": "Server error"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
			return
		}
	} else if user == nil {
		message := map[string]string{"message": "User Not Found"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		message := map[string]string{"message": "Username atau password salah"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return
	}

	expTimeToken := time.Now().Add(time.Hour * 24)

	claims := &config.JWTClaim{
		UserId:   int(user.ID),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-postgres",
			ExpiresAt: jwt.NewNumericDate(expTimeToken),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(JWT_KEY)
	if err != nil {
		message := map[string]string{"message": "Server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})
	message := map[string]string{"message": "success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		message := map[string]string{"message": "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}

	defer r.Body.Close()

	if user, err := models.GetUserByUsername(userInput.Username); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := map[string]string{"message": "Server error"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
			return
		}
	} else if user != nil {
		message := map[string]string{"message": "User Already registered"}
		utils.SendJSONResponse(w, http.StatusConflict, message)
		return
	}

	if hashedPassword, err := utils.HashPassword(userInput.Password); err != nil {
		message := map[string]string{"message": "Internal Server Error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	} else {
		userInput.Password = *hashedPassword
	}

	if err := models.CreateUser(&userInput); err != nil {
		message := map[string]string{"message": "Failed to Create User"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]string{"message": "success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	message := map[string]string{"message": "Logout Success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}
