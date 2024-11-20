package usecase

import (
	"baf-credit-score/model"
	"baf-credit-score/model/dto"
	"baf-credit-score/tests/usecase/mock"
	"baf-credit-score/usecase"
	"errors"
	"testing"
	"time"

	testifyMock "github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/suite"
)

type CustomerUsecaseTestSuite struct {
	suite.Suite
	mockRepo   *mock.MockCustomerRepository
	customerUc usecase.CustomerUsecase
}

func (suite *CustomerUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mock.MockCustomerRepository)
	suite.customerUc = usecase.NewCustomerUsecase(suite.mockRepo)
}

func TestCustomerUsecase(t *testing.T) {
	suite.Run(t, new(CustomerUsecaseTestSuite))
}

func (suite *CustomerUsecaseTestSuite) TestRegisterCustomer_Success() {
	// Arrange
	payload := dto.CustomerRequestDto{
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   "02-01-2006",
	}

	suite.mockRepo.On("Save", testifyMock.AnythingOfType("model.Customer")).Return(nil)

	// Act
	err := suite.customerUc.RegisterCustomer(payload)

	// Assert
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *CustomerUsecaseTestSuite) TestRegisterCustomer_InvalidDate() {
	// Arrange
	payload := dto.CustomerRequestDto{
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   "invalid-date",
	}

	// Act
	err := suite.customerUc.RegisterCustomer(payload)

	// Assert
	suite.Error(err)
	suite.Contains(err.Error(), "failed to parse date")
}

func (suite *CustomerUsecaseTestSuite) TestFindCustomerById_Success() {
	// Arrange
	expectedCustomer := model.Customer{
		BaseModel: model.BaseModel{
			ID:        "test-id",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   time.Now(),
	}

	suite.mockRepo.On("Get", "test-id").Return(expectedCustomer, nil)

	// Act
	customer, err := suite.customerUc.FindCustomerById("test-id")

	// Assert
	suite.NoError(err)
	suite.Equal(expectedCustomer, customer)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *CustomerUsecaseTestSuite) TestFindCustomerById_NotFound() {
	// Arrange
	suite.mockRepo.On("Get", "non-existent-id").
		Return(model.Customer{}, errors.New("customer not found"))

	// Act
	customer, err := suite.customerUc.FindCustomerById("non-existent-id")

	// Assert
	suite.Error(err)
	suite.Equal(model.Customer{}, customer)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *CustomerUsecaseTestSuite) TestFindAll_Success() {
	// Arrange
	expectedCustomers := []model.Customer{
		{
			BaseModel: model.BaseModel{
				ID:        "test-id-1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			FullName:    "John Doe",
			PhoneNumber: "081234567890",
			NIK:         "1234567890123456",
			Address:     "Test Address 1",
			Status:      "active",
			BirthDate:   time.Now(),
		},
		{
			BaseModel: model.BaseModel{
				ID:        "test-id-2",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			FullName:    "Jane Doe",
			PhoneNumber: "089876543210",
			NIK:         "6543210987654321",
			Address:     "Test Address 2",
			Status:      "active",
			BirthDate:   time.Now(),
		},
	}

	suite.mockRepo.On("List").Return(expectedCustomers, nil)

	// Act
	customers, err := suite.customerUc.FindAll()

	// Assert
	suite.NoError(err)
	suite.Equal(expectedCustomers, customers)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *CustomerUsecaseTestSuite) TestUpdateCustomer_Success() {
	// Arrange
	payload := dto.CustomerRequestDto{
		Id:          "test-id",
		FullName:    "John Doe Updated",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address Updated",
		Status:      "inactive",
		BirthDate:   "02-01-2006",
	}

	suite.mockRepo.On("Update", testifyMock.AnythingOfType("model.Customer")).Return(nil)

	// Act
	err := suite.customerUc.UpdateCustomer(payload)

	// Assert
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite * CustomerUsecaseTestSuite) TestUpdateCustomer_Failed(){
	// Arrange
	payload := dto.CustomerRequestDto{
		Id:          "test-id",
		FullName:    "John Doe Updated",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address Updated",
		Status:      "inactive",
		BirthDate:   "02-01-2006",
	}

	suite.mockRepo.On("Update", testifyMock.AnythingOfType("model.Customer")).Return(errors.New("Error nih"))

	// Act
	err := suite.customerUc.UpdateCustomer(payload)

	// Assert
	suite.Error(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *CustomerUsecaseTestSuite) TestUpdateCustomer_InvalidDate() {
	// Arrange
	payload := dto.CustomerRequestDto{
		Id: "test-id",
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   "invalid-date",
	}

	// Act
	err := suite.customerUc.UpdateCustomer(payload)

	// Assert
	suite.Error(err)
	suite.Contains(err.Error(), "failed to parse date")
}

func (suite *CustomerUsecaseTestSuite) TestDeleteCustomer_Success() {
	// Arrange
	customerID := "test-id"
	mockCustomer := model.Customer{
		BaseModel: model.BaseModel{ID: customerID},
	}

	suite.mockRepo.On("Get", customerID).Return(mockCustomer, nil)
	suite.mockRepo.On("Delete", customerID).Return(nil)

	// Act
	err := suite.customerUc.DeleteCustomer(customerID)

	// Assert
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *CustomerUsecaseTestSuite) TestDeleteCustomer_NotFound() {
	// Arrange
	customerID := "non-existent-id"

	suite.mockRepo.On("Get", customerID).
		Return(model.Customer{}, errors.New("customer not found"))

	// Act
	err := suite.customerUc.DeleteCustomer(customerID)

	// Assert
	suite.Error(err)
	suite.mockRepo.AssertExpectations(suite.T())
}
