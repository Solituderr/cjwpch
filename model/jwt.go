package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
