package usecase

import (
	"baf-credit-score/model/dto"
	"baf-credit-score/utils/service"
)

type AuthenticationUsecase interface {
	Login(email string, password string) (string, error)
	Register(payload dto.UserRequestDto) error
}

type authenticationUsecase struct {
	userUc     UserUsecase
	jwtService service.JwtService
}

// Login implements AuthenticationUsecase.
func (a *authenticationUsecase) Login(email string, password string) (string, error) {
	user, err := a.userUc.FindByEmailPassword(email,password)
	if err != nil {
		return "",err
	}
	token, errGenerateToken := a.jwtService.CreateAccessToken(user)
	if errGenerateToken != nil {
		return "",errGenerateToken
	}
	return token,nil
}

// Register implements AuthenticationUsecase.
func (a *authenticationUsecase) Register(payload dto.UserRequestDto) error {
	return a.userUc.RegisterUser(payload)
}

func NewAuthenticationUsecase(userUsecase UserUsecase, jwtService service.JwtService) AuthenticationUsecase {
	return &authenticationUsecase{
		userUc:     userUsecase,
		jwtService: jwtService,
	}
}
