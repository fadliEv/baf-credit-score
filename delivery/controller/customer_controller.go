package controller

import (
	"baf-credit-score/model"
	"baf-credit-score/usecase"
)

type CustomerController interface {
	RegisterCustomer(newCustomer model.Customer) error
}

type customerController struct {
	uc usecase.CustomerUsecase
}

// RegisterCustomer implements CustomerController.
func (c *customerController) RegisterCustomer(newCustomer model.Customer) error {
	return c.uc.RegisterCustomer(newCustomer)
}

func NewCustomerController(usecase usecase.CustomerUsecase) CustomerController {
	return &customerController{
		uc: usecase,
	}
}
