package config

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	UserId   int
	Username string
	jwt.RegisteredClaims
}

type AdminJWTClaim struct {
	AdminId  int
	RoleID   int
	Username string
	jwt.RegisteredClaims
}
