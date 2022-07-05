package domain

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Addres   string `json:"address"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginReponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type JWTClaim struct {
	Username string `json:"username"`
	Id       int64  `json:"Id"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type TokenMap struct {
	U *User
	T *LoginReponse
}
