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
	UserId    int  `json:"user_id"`
	IsAdmin   bool `json:"is_admin"`
	IsManager bool `json:"is_manager"`
	IsDoctor  bool `json:"is_doctor"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId    int  `json:"user_id"`
	IsAdmin   bool `json:"is_admin"`
	IsManager bool `json:"is_manager"`
	IsDoctor  bool `json:"is_doctor"`
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

type AdminAccountResponse struct {
	Id        int      `json:"id"`
	LastName  string   `json:"lastName"`
	FirstName string   `json:"firstName"`
	Username  string   `json:"username"`
	Roles     []string `json:"roles"`
	IsDeleted bool     `json:"isDeleted"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type HospitalResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	ContactPhone string `json:"contactPhone"`
	// Rooms        []string `json:"rooms"`
}

type HospitalRequest struct {
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	ContactPhone string   `json:"contactPhone"`
	Rooms        []string `json:"rooms"`
}
