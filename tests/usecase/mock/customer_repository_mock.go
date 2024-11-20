package mock

import (
	"baf-credit-score/model"	
	"github.com/stretchr/testify/mock"
)


type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) Save(payload model.Customer) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *MockCustomerRepository) Get(id string) (model.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(model.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Update(payload model.Customer) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *MockCustomerRepository) List() ([]model.Customer, error) {
	args := m.Called()
	return args.Get(0).([]model.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}