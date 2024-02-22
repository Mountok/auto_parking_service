package models




type Car struct {
	Id int64 `json:"id"`
	Brand string `json:"brand"`
	Colour string `json:"colour"`
	LicensePlate string `json:"license_plate"`
	Owner User `json:"owner"`
}