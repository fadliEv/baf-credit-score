package dto

type CustomerRequestDto struct {
	Id          string `json:"id"`
	FullName    string `json:"fullName" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	NIK         string `json:"nik" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Status      string `json:"status" binding:"required"`
	BirthDate   string `json:"birthDate" binding:"required,datetime=02-01-2006"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type CustomerResponseDto struct {
	BaseModelResponseDto
	FullName    string              `json:"fullName"`
	PhoneNumber string              `json:"phoneNumber"`
	NIK         string              `json:"nik"`
	Address     string              `json:"address"`
	Status      string              `json:"status"`
	BirthDate   string              `json:"birthDate"`
	User        *UserResponseDto     `json:"user,omitempty"`
	Credits     []CreditResponseDto `json:"credits,omitempty"`
}
