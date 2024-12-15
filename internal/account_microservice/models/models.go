package models

import "github.com/dgrijalva/jwt-go"

type SignUpRequest struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenInfo struct {
	UserId  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}

type AccountUpdate struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Password  string `json:"password"`
}

type AccountResponse struct {
	Id        int      `json:"id"`
	LastName  string   `json:"lastName"`
	FirstName string   `json:"firstName"`
	Username  string   `json:"username"`
	Roles     []string `json:"roles"`
}

type DoctorResponse struct {
	Id        int    `json:"id"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Username  string `json:"username"`
}

type AdminAccountRequest struct {
	LastName  string   `json:"lastName"`
	FirstName string   `json:"firstName"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Roles     []string `json:"roles"`
}
