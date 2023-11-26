package middlewares

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rakhazufar/go-project/pkg/config"
	"github.com/rakhazufar/go-project/pkg/utils"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response:= map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			}
		}

		tokenString := c.Value

		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v,_ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response:= map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response:= map[string]string{"message": "Unauthorized, Token Expired!"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			default:
				response:= map[string]string{"message": "Unauthorized"}
				utils.SendJSONResponse(w, http.StatusUnauthorized, response)
				return
			}	  
		}

		if !token.Valid {
			response:= map[string]string{"message": "Unauthorized"}
			utils.SendJSONResponse(w, http.StatusUnauthorized, response)
			return
		}
		
		ctx := context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "userId", claims.UserId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}