package usecase

import (
	"baf-credit-score/model"
	"baf-credit-score/model/dto"
	"baf-credit-score/repository"
	"baf-credit-score/utils"
)

type CustomerUsecase interface {
	RegisterCustomer(payload dto.CustomerRequestDto) error
	FindCustomerById(id string) (model.Customer, error)
}

type customerUsecase struct {
	repo repository.CustomerRepository
}

// FindCustomerById implements CustomerUsecase.
func (c *customerUsecase) FindCustomerById(id string) (model.Customer, error) {
	return c.repo.FindCustomerById(id)
}

// RegisterCustomer implements CustomerUsecase.
func (c *customerUsecase) RegisterCustomer(payload dto.CustomerRequestDto) error {
	parseBirthDate, err := utils.ParseDate(payload.BirthDate)
	if err != nil {
		return err
	}
	customer := model.Customer{
		FullName:    payload.FullName,
		PhoneNumber: payload.PhoneNumber,
		Status:      payload.Status,
		NIK:         payload.NIK,
		Address:     payload.Address,
		BirthDate:   parseBirthDate,
	}
	return c.repo.InsertCustomer(customer)
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		repo: repo,
	}
}
