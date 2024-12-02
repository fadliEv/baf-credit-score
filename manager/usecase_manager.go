package manager

import "baf-credit-score/usecase"

type UsecaseManager interface {
	AuthUsecase() usecase.AuthenticationUsecase
	CreditUsecase() usecase.CreditUsecase
	CreditScoreUsecase() usecase.CreditScoreUsecase
	CustomerUsecase() usecase.CustomerUsecase
	UserUsecas() usecase.UserUsecase
}

type usecaseManager struct {
	repo RepositoryManager
}

// AuthUsecase implements UsecaseManager.
func (u *usecaseManager) AuthUsecase() usecase.AuthenticationUsecase {
	return usecase.NewAuthenticationUsecase(u.UserUsecas(),u.repo.JwtService())
}

// CreditScoreUsecase implements UsecaseManager.
func (u *usecaseManager) CreditScoreUsecase() usecase.CreditScoreUsecase {
	return usecase.NewCreditScoreUsecase(u.repo.CreditScoreRepo())
}

// CreditUsecase implements UsecaseManager.
func (u *usecaseManager) CreditUsecase() usecase.CreditUsecase {
	return usecase.NewCreditUsecase(u.repo.CreditRepo(),u.CreditScoreUsecase())
}

// CustomerUsecase implements UsecaseManager.
func (u *usecaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(u.repo.CustomerRepo())
}

// UserUsecas implements UsecaseManager.
func (u *usecaseManager) UserUsecas() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.repo.UserRepo())
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repo: repoManager,
	}
}
