package controller

import (
	"baf-credit-score/delivery/controller"
	"baf-credit-score/model"
	"baf-credit-score/model/dto"
	"baf-credit-score/tests/controller/mock"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerControllerTestSuite struct {
	suite.Suite
	router     *gin.Engine
	customerUc *mock.MockCustomerUsecase
	authMid    *mock.MockAuthMiddleware
	custCtrl   *controller.CustomerController
}

func (s *CustomerControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.router = gin.New()
	s.customerUc = new(mock.MockCustomerUsecase)
	s.authMid = new(mock.MockAuthMiddleware)

	rg := s.router.Group("/api/v1")
	s.custCtrl = controller.NewCustomerController(s.customerUc, rg, s.authMid)
	s.custCtrl.Route()
}

func TestCustomerController(t *testing.T) {
	suite.Run(t, new(CustomerControllerTestSuite))
}

func (s *CustomerControllerTestSuite) TestCreateHandler_Success() {
	// Arrange
	customerPayload := dto.CustomerRequestDto{
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   "02-01-2006",
	}

	s.customerUc.On("RegisterCustomer", customerPayload).Return(nil)

	jsonPayload, err := json.Marshal(customerPayload)
	s.NoError(err)

	// Create request
	req, err := http.NewRequest(http.MethodPost, "/api/v1/customers", bytes.NewBuffer(jsonPayload))
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.NoError(err)

	status := response["status"].(map[string]interface{})
	s.Equal(float64(http.StatusOK), status["code"])
	s.Equal("Success Register Customer", status["description"])

	s.customerUc.AssertExpectations(s.T())
}

func (s *CustomerControllerTestSuite) TestCreateHandler_InvalidBinding() {
	// Arrange
	customerPayload := dto.CustomerRequestDto{
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   "2000-01-01",
	}

	s.customerUc.On("RegisterCustomer", customerPayload).Return(nil)

	jsonPayload, err := json.Marshal(customerPayload)
	s.NoError(err)

	// Create request
	req, err := http.NewRequest(http.MethodPost, "/api/v1/customers", bytes.NewBuffer(jsonPayload))
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusBadRequest, recorder.Code)
	assert.Contains(s.T(), response["description"], "validation for 'BirthDate' failed on the 'datetime' tag")
}

func (s *CustomerControllerTestSuite) TestCreateHandler_Failed() {
	// Arrange
	customerPayload := dto.CustomerRequestDto{
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   "01-01-2001",
	}

	s.customerUc.On("RegisterCustomer", customerPayload).Return(errors.New("ada error dari usecase"))

	jsonPayload, err := json.Marshal(customerPayload)
	s.NoError(err)

	// Create request
	req, err := http.NewRequest(http.MethodPost, "/api/v1/customers", bytes.NewBuffer(jsonPayload))
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusInternalServerError, recorder.Code)
	assert.Contains(s.T(),response["description"],"ada error dari usecase")
}

func (s *CustomerControllerTestSuite) TestListHandler_Success() {
	// Arrange
	customers := []model.Customer{
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
	}

	s.customerUc.On("FindAll").Return(customers, nil)

	// Create request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/customers", nil)
	s.NoError(err)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.NoError(err)

	status := response["status"].(map[string]interface{})
	s.Equal(float64(http.StatusOK), status["code"])
	s.Equal("Success Get All Customer", status["description"])

	s.customerUc.AssertExpectations(s.T())
}

func (s *CustomerControllerTestSuite) TestListHandler_Failed() {
	s.customerUc.On("FindAll").Return([]model.Customer{},errors.New("error bro"))

	// Create request
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/customers", nil)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	
	assert.Equal(s.T(),recorder.Code,http.StatusInternalServerError)
}

func (s *CustomerControllerTestSuite) TestFindByIdHandler_Success() {
	// Arrange
	customerID := "test-id"
	expectedCustomer := model.Customer{
		BaseModel: model.BaseModel{
			ID:        customerID,
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

	s.customerUc.On("FindCustomerById", customerID).Return(expectedCustomer, nil)

	// Create request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/customers/"+customerID, nil)
	s.NoError(err)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.NoError(err)

	status := response["status"].(map[string]interface{})
	s.Equal(float64(http.StatusOK), status["code"])
	s.Equal("Success Get Customer By Id", status["description"])

	s.customerUc.AssertExpectations(s.T())
}

func (s *CustomerControllerTestSuite) TestFindByIdHandler_NotFound() {
	// Arrange
	customerID := "non-existent-id"

	s.customerUc.On("FindCustomerById", customerID).
		Return(model.Customer{}, errors.New("customer not found"))

	// Create request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/customers/"+customerID, nil)
	s.NoError(err)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	// Assert
	s.Equal(http.StatusInternalServerError, recorder.Code)

	s.customerUc.AssertExpectations(s.T())
}

func (s *CustomerControllerTestSuite) TestDeleteHandler_Success() {
	// Arrange
	customerID := "test-id"

	s.customerUc.On("DeleteCustomer", customerID).Return(nil)

	// Create request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/customers/"+customerID, nil)
	s.NoError(err)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	// Assert
	s.Equal(http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.NoError(err)

	status := response["status"].(map[string]interface{})
	s.Equal(float64(http.StatusOK), status["code"])
	s.Equal("Success Delete Customer", status["description"])

	s.customerUc.AssertExpectations(s.T())
}

func (s *CustomerControllerTestSuite) TestDeleteHandler_Failed() {
	// Arrange
	customerID := "test-id"

	s.customerUc.On("DeleteCustomer", customerID).Return(errors.New("data is not found"))

	// Create request
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/customers/"+customerID, nil)
	s.NoError(err)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.NoError(err)
	assert.Equal(s.T(), http.StatusInternalServerError, recorder.Code)	
}

func (s *CustomerControllerTestSuite) TestUpdateByIdHandler_Success() {
	// Arrange
	customerPayload := dto.CustomerRequestDto{
		Id: "test-id",
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   "01-01-2001",
	}

	s.customerUc.On("UpdateCustomer", customerPayload).Return(nil)

	jsonPayload, err := json.Marshal(customerPayload)
	s.NoError(err)

	// Create request
	req, err := http.NewRequest(http.MethodPut, "/api/v1/customers", bytes.NewBuffer(jsonPayload))
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, recorder.Code)	
}

func (s *CustomerControllerTestSuite) TestUpdateByIdHandler_InvalidBinding() {
	payload := map[string]string{
		"Id":          "test-id",
		"FullName":    "John Doe",
		"PhoneNumber": "081234567890",
		"NIK":         "1234567890123456",
	}

	s.customerUc.On("UpdateCustomer", payload).Return(nil)

	jsonPayload, err := json.Marshal(payload)
	s.NoError(err)

	// Create request
	req, err := http.NewRequest(http.MethodPut, "/api/v1/customers", bytes.NewBuffer(jsonPayload))
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusBadRequest, recorder.Code)	
}

func (s *CustomerControllerTestSuite) TestUpdateByIdHandler_Failed() {
	customerPayload := dto.CustomerRequestDto{
		Id: "test-id",
		FullName:    "John Doe",
		PhoneNumber: "081234567890",
		NIK:         "1234567890123456",
		Address:     "Test Address",
		Status:      "active",
		BirthDate:   "01-01-2001",
	}

	s.customerUc.On("UpdateCustomer", customerPayload).Return(errors.New("Error update"))

	jsonPayload, err := json.Marshal(customerPayload)
	s.NoError(err)

	// Create request
	req, err := http.NewRequest(http.MethodPut, "/api/v1/customers", bytes.NewBuffer(jsonPayload))
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	s.router.ServeHTTP(recorder, req)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusInternalServerError, recorder.Code)	
}