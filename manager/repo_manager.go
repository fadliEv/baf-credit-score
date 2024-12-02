package manager

import (
	"baf-credit-score/repository"
	"baf-credit-score/utils/service"
)

type RepositoryManager interface {
	CreditRepo() repository.CreditRepository
	CreditScoreRepo() repository.CreditScoreRepository
	CustomerRepo() repository.CustomerRepository
	UserRepo() repository.UserRepository
	JwtService() service.JwtService
}

type repositoryManager struct {
	infra InfraManager
}

// JwtService implements RepositoryManager.
func (r *repositoryManager) JwtService() service.JwtService {
	return service.NewJwtService(r.infra.Config().TokenConfig)
}

// CreditRepo implements RepositoryManager.
func (r *repositoryManager) CreditRepo() repository.CreditRepository {
	return repository.NewCreditRepository(r.infra.Conn())
}

// CreditScoreRepo implements RepositoryManager.
func (r *repositoryManager) CreditScoreRepo() repository.CreditScoreRepository {
	return repository.NewCreditScoreRepository(r.infra.Conn())
}

// CustomerRepo implements RepositoryManager.
func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

// UserRepo implements RepositoryManager.
func (r *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
