package models

import "github.com/dgrijalva/jwt-go"


type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}

func (t *Token)GenerateToken() {
	
}