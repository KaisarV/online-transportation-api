package entity

type User struct {
	Id               int    `json:"id"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone"`
	Password         string `json:"password,omitempty"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Gender           string `json:"gender"`
	Province         string `json:"province"`
	City             string `json:"city"`
	Address          string `json:"address"`
	RegistrationDate string `json:"registration_date"`
	IsVerified       int    `json:"is_verified,omitempty"`
}
