package models

import (
	"github.com/dgrijalva/jwt-go"
)

type FormRegister struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
	FullName string `json:"full_name" valid:"required"`
	Picture  string `json:"picture"`
}

type FormAuth struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type MyClaims struct {
	jwt.StandardClaims
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type AuthResult struct {
	Token string `json:"token"`
}

type UserJWT struct {
	ID       int64
	Username string
}
