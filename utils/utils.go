package utils

import "github.com/golang-jwt/jwt/v4"

type Request struct {
	Status int
	Body   interface{}
}
type AuthRequest struct {
	Email    string
	Password string
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
