package models

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
