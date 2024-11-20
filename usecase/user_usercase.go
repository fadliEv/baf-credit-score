package usecase

import (
	"baf-credit-score/model"
	"baf-credit-score/model/dto"
	"baf-credit-score/repository"
	"baf-credit-score/utils/common"
	"baf-credit-score/utils/payload"
)

type UserUsecase interface {
	RegisterUser(payload dto.UserRequestDto) (dto.UserResponseDto, error)
	FindUserById(id string) (model.User, error)
	FindUserByEmail(email string) (model.User, error)
	FindByEmailPassword(email string, password string) (model.User, error)
	FindAll(size int, page int) ([]dto.UserResponseDto, payload.Paging, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

// FindByEmailPassword implements UserUsecase.
func (u *userUsecase) FindByEmailPassword(email string, password string) (model.User, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	errCheckPass := common.CheckPasswordHash(password, user.Password)
	if errCheckPass != nil {
		return model.User{}, errCheckPass
	}
	return user, nil
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
func (u *userUsecase) RegisterUser(payload dto.UserRequestDto) (dto.UserResponseDto, error) {
	hashedPass, err := common.HashPassword(payload.Password)
	if err != nil {
		return dto.UserResponseDto{}, err
	}
	user := model.User{
		Email:    payload.Email,
		Password: hashedPass,
		Role:     payload.Role,
	}
	errSave := u.repo.Save(user)
	if errSave != nil {
		return dto.UserResponseDto{}, errSave
	}
	userResponse := dto.UserResponseDto{
		Email: payload.Email,
		Role:  payload.Role,
	}
	return userResponse, nil

}

func (u *userUsecase) FindAll(size int, page int) ([]dto.UserResponseDto, payload.Paging, error) {
	totalRecords, err := u.repo.GetTotal()
	if err != nil {
		return nil, payload.Paging{}, err
	}
	totalPages := (int(totalRecords) + size - 1) / size
	offset := (page - 1) * size
	paging := payload.Paging{
		Page:        page,
		TotalRows:   totalRecords,
		RowsPerPage: size,
		TotalPages:  totalPages,
	}
	users, err := u.repo.List(size, offset)
	var userResponses []dto.UserResponseDto
	for _, user := range users {
		userResponses = append(userResponses, u.mappingResponse(user))
	}
	if err != nil {
		return nil, payload.Paging{}, err
	}
	return userResponses, paging, nil
}

func (u *userUsecase) mappingResponse(payload model.User) dto.UserResponseDto {
	return dto.UserResponseDto{
		BaseModelResponseDto: dto.BaseModelResponseDto{
			Id:        payload.ID,
			CreatedAt: payload.CreatedAt,
			UpdatedAt: payload.UpdatedAt,
		},
		Email: payload.Email,
		Role:  payload.Role,
	}
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}
