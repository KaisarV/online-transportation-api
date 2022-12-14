package entity

type Driver struct {
	Id               int    `json:"id"`
	UserId           string `json:"email"`
	NIK              string `json:"phone"`
	VehicleType      string `json:"password,omitempty"`
	VehicleBrand     string `json:"firstname"`
	BrandType        string `json:"brand_type"`
	LisencePlate     string `json:"lastname"`
	IsVaccinated     string `json:"is_vaccinated"`
	IsVerified       int    `json:"is_verified,omitempty"`
	RegistrationDate string `json:"registration_date"`
}
