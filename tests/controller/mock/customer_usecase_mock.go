package mock

import (
	"baf-credit-score/model"
	"baf-credit-score/model/dto"	
	"github.com/stretchr/testify/mock"
)

type MockCustomerUsecase struct {
	mock.Mock
}

func(m *MockCustomerUsecase) RegisterCustomer(payload dto.CustomerRequestDto) error {
	args := m.Called(payload)
	return args.Error(0)
}
func(m *MockCustomerUsecase) FindCustomerById(id string) (model.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(model.Customer), args.Error(1)
}
func(m *MockCustomerUsecase) FindAll() ([]model.Customer, error) {
	args := m.Called()
	return args.Get(0).([]model.Customer), args.Error(1)
}
func(m *MockCustomerUsecase) UpdateCustomer(payload dto.CustomerRequestDto) error {
	args := m.Called(payload)
	return args.Error(0)
}
func(m *MockCustomerUsecase) DeleteCustomer(id string) error {
	args := m.Called(id)
	return args.Error(0)
}