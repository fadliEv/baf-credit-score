package dto

type UserRequestDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserResponseDto struct {
	BaseModelResponseDto
	Email    string `json:"email"`	
	Role     string `json:"role"`
}
