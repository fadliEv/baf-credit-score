package usecase

import (
	"baf-credit-score/model"
	"baf-credit-score/model/dto"
	"baf-credit-score/repository"
	"baf-credit-score/utils/common"
)

type UserUsecase interface {
	RegisterUser(payload dto.UserRequestDto) (dto.UserResponseDto,error)
	FindUserById(id string) (model.User, error)
	FindUserByEmail(email string) (model.User, error)
	FindByEmailPassword(email string, password string) (model.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

// FindByEmailPassword implements UserUsecase.
func (u *userUsecase) FindByEmailPassword(email string, password string) (model.User, error) {
	user,err := u.repo.GetByEmail(email)
	if err != nil {
		return model.User{},err
	}
	errCheckPass := common.CheckPasswordHash(password,user.Password)
	if errCheckPass != nil {
		return model.User{},errCheckPass
	}
	return user,nil
}

// FindUserByEmail implements UserUsecase.
func (u *userUsecase) FindUserByEmail(email string) (model.User, error) {
	return u.repo.GetByEmail(email)
}

// FindUserById implements UserUsecase.
func (u *userUsecase) FindUserById(id string) (model.User, error) {
	return u.repo.Get(id)
}

// RegisterUser implements UserUsecase.
func (u *userUsecase) RegisterUser(payload dto.UserRequestDto) (dto.UserResponseDto,error) {
	hashedPass, err := common.HashPassword(payload.Password)
	if err != nil {
		return dto.UserResponseDto{},err
	}
	user := model.User{
		Email:    payload.Email,
		Password: hashedPass,
		Role:     payload.Role,
	}
	userResponse := dto.UserResponseDto {
		Email: payload.Email,
		Role: payload.Role,
	}
	errSave := u.repo.Save(user)
	if errSave != nil {
		return dto.UserResponseDto{},err
	}
	return userResponse,nil

}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}
