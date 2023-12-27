package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/rakhazufar/go-project/pkg/config"
	"github.com/rakhazufar/go-project/pkg/utils"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			}
		}

		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
		var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

		tokenString := c.Value

		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{"message": "Unauthorized, Token Expired!"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			utils.SendJSONResponse(w, http.StatusUnauthorized, response)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "userId", claims.UserId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func AdminJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.SendJSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
		var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

		claims := &config.AdminJWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{"message": "Unauthorized, Token Expired!"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			utils.SendJSONResponse(w, http.StatusUnauthorized, response)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "adminId", claims.AdminId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func AdministratorJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.SendJSONResponse(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
		var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

		claims := &config.AdminJWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{"message": "Unauthorized, Token Expired!"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			utils.SendJSONResponse(w, http.StatusUnauthorized, response)
			return
		}
		fmt.Println(claims.RoleID)
		if claims.RoleID != 1 {
			response := map[string]string{"message": "Unauthentication"}
			utils.SendJSONResponse(w, http.StatusUnauthorized, response)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "adminId", claims.AdminId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
