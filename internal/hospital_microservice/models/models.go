package models

import "github.com/dgrijalva/jwt-go"

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

type TokenClaims struct {
	jwt.StandardClaims
	UserId    int  `json:"user_id"`
	IsAdmin   bool `json:"is_admin"`
	IsManager bool `json:"is_manager"`
	IsDoctor  bool `json:"is_doctor"`
}
