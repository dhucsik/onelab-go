package models

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	Username string
	UserRole string
	jwt.StandardClaims
}

type ContextUserData struct {
	Usename  string
	UserRole string
}

type contextKey string

const ContextKey = contextKey("UserData")
