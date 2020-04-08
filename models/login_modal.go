package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Register struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type RegisterResponse struct {
	StatusCode int
	Body       map[string]interface{}
}
