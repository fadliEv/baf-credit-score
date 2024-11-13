package usecase

import (
	"baf-credit-score/model"
	"baf-credit-score/repository"
)

type CustomerUsecase interface {
	RegisterCustomer(newCustomer model.Customer) error
}

type customerUsecase struct {
	repo repository.CustomerRepository
}

// RegisterCustomer implements CustomerUsecase.
func (c *customerUsecase) RegisterCustomer(newCustomer model.Customer) error {
	return c.repo.InsertCustomer(newCustomer)
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		repo: repo,
	}
}
