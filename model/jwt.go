package model

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
