package tokencontrollers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/rakhazufar/go-project/pkg/utils"
)

type Token struct {
	Token string `json:"token"`
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

	secretKey := JWT_KEY

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		utils.SendJSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	defer r.Body.Close()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi algoritma token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		utils.SendJSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Invalid token", "error": err.Error()})
		return
	}
	// Cek validitas token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		adminId, ok := claims["Username"]
		if !ok {
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Username claim missing"})
			return
		}

		adminIdStr, ok := adminId.(string)
		if !ok {
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Error converting Username to string"})
			return
		}

		utils.SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Token is valid", "Username": adminIdStr})
	} else {
		utils.SendJSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}
}
