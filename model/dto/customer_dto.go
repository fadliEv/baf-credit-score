package dto

type CustomerRequestDto struct {
	Id          string `json:"id"`
	FullName    string `json:"fullName" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	NIK         string `json:"nik" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Status      string `json:"status" binding:"required"`
	BirthDate   string `json:"birthDate" binding:"required,datetime=02-01-2006"`
}