package tokencontrollers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/rakhazufar/go-project/pkg/utils"
)

type Token struct {
	Token string `json:"token"`
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	var tokenInput Token

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

	secretKey := JWT_KEY

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tokenInput); err != nil {
		message := map[string]string{"message": "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}

	fmt.Println("Secret key: ", string(secretKey))
	fmt.Println("Received token: ", tokenInput.Token)

	defer r.Body.Close()

	token, err := jwt.Parse(tokenInput.Token, func(token *jwt.Token) (interface{}, error) {
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
		fmt.Println(adminIdStr)
		if !ok {
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Error converting Username to string"})
			return
		}

		utils.SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Token is valid", "Username": adminIdStr})
	} else {
		utils.SendJSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}
}
