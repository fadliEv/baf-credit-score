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
	FindAll() ([]model.Customer, error)
	UpdateCustomer(payloada dto.CustomerRequestDto) error
	DeleteCustomer(id string) error
}

type customerUsecase struct {
	repo repository.CustomerRepository
}

// DeleteCustomer implements CustomerUsecase.
func (c *customerUsecase) DeleteCustomer(id string) error {
	_, err := c.FindCustomerById(id) // Cek terlebih dahulu jika id tidak ada, karena secara default delete record pada postgresql tidak ada pengecekan id
	if err != nil {
		return err
	}
	return c.repo.Delete(id)
}

// FindAll implements CustomerUsecase.
func (c *customerUsecase) FindAll() ([]model.Customer, error) {
	return c.repo.List()
}

// UpdateCustomer implements CustomerUsecase.
func (c *customerUsecase) UpdateCustomer(payload dto.CustomerRequestDto) error {
	customerModel, err := c.mappingToModel(payload)
	if err != nil {
		return err
	}
	customerModel.ID = payload.Id
	return c.repo.Update(customerModel)
}

// FindCustomerById implements CustomerUsecase.
func (c *customerUsecase) FindCustomerById(id string) (model.Customer, error) {
	return c.repo.Get(id)
}

// RegisterCustomer implements CustomerUsecase.
func (c *customerUsecase) RegisterCustomer(payload dto.CustomerRequestDto) error {
	customerModel, err := c.mappingToModel(payload)
	if err != nil {
		return err
	}
	return c.repo.Save(customerModel)
}

func (c *customerUsecase) mappingToModel(payload dto.CustomerRequestDto) (model.Customer,error){
	parseBirthDate, err := utils.ParseDate(payload.BirthDate)
	if err != nil {
		return model.Customer{}, err
	}
	return model.Customer{		 
		 FullName: payload.FullName,
		 PhoneNumber: payload.PhoneNumber,
		 Status: payload.Status,
		 NIK: payload.NIK,
		 Address: payload.Address,
		 BirthDate: parseBirthDate,		 
	},nil
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		repo: repo,
	}
}
