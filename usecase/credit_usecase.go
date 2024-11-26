package usecase

import (
	"baf-credit-score/model"
	"baf-credit-score/model/dto"
	"baf-credit-score/repository"
	"baf-credit-score/utils/common"
	"baf-credit-score/utils/payload"
)

type CreditUsecase interface {
	CreateCredit(payload dto.CreditRequestDto) error
	GetCreditByID(id string) (model.Credit, error)
	GetCreditsByCustomer(customerID string) (dto.CustomerCreditResponseDto, error)
	FindAll(size int, page int) ([]dto.CreditResponseDto, payload.Paging, error)
}

type creditUsecase struct {
	repo repository.CreditRepository
}

func (c *creditUsecase) CreateCredit(payload dto.CreditRequestDto) error {
	creditModel, err := c.mappingToModel(payload)
	if err != nil {
		return err
	}
	return c.repo.Save(creditModel)
}

func (c *creditUsecase) GetCreditByID(id string) (model.Credit, error) {
	return c.repo.Get(id)
}

func (c *creditUsecase) GetCreditsByCustomer(customerID string) (dto.CustomerCreditResponseDto, error) {	
	credits, err := c.repo.GetByCustomer(customerID)
	if err != nil {
		return dto.CustomerCreditResponseDto{},err
	}
	customerCredits := dto.CustomerCreditResponseDto{
		Customer: c.mapCustomerToResponse(credits[0].Customer),
	}
	for _, credit := range credits {		
		creditResponse := c.mappingToResponse(credit)
		creditResponse.Customer = nil // tidak perlu ditampilkan didalam response
		customerCredits.Credits = append(customerCredits.Credits, creditResponse)
	}
	return customerCredits,nil
}

func (c *creditUsecase) UpdateCredit(payload dto.CreditRequestDto) error {
	creditModel, err := c.mappingToModel(payload)
	if err != nil {
		return err
	}
	creditModel.ID = payload.Id
	return c.repo.Update(creditModel)
}

func (c *creditUsecase) mappingToModel(payload dto.CreditRequestDto) (model.Credit, error) {
	return model.Credit{
		CustomerID:       payload.CustomerID,
		AppNumber:        payload.AppNumber,
		ProductType:      payload.ProductType,
		LoanAmount:       payload.LoanAmount,
		Tenure:           payload.Tenure,
		EmploymentStatus: payload.EmploymentStatus,
		MonthlyIncome:    payload.MonthlyIncome,
		Status:           payload.Status,
		RejectionReason:  payload.RejectionReason,
	}, nil
}

func (c *creditUsecase) mappingToResponse(payload model.Credit) dto.CreditResponseDto {
	return dto.CreditResponseDto{
		BaseModelResponseDto: dto.BaseModelResponseDto{
			Id:        payload.BaseModel.ID,
			CreatedAt: payload.BaseModel.CreatedAt,
			UpdatedAt: payload.BaseModel.UpdatedAt,
		},
		AppNumber: payload.AppNumber,
		Customer: &dto.CustomerResponseDto{
			BaseModelResponseDto: dto.BaseModelResponseDto{
				Id:        payload.Customer.BaseModel.ID,
				CreatedAt: payload.Customer.CreatedAt,
				UpdatedAt: payload.Customer.UpdatedAt,
			},
			FullName:    payload.Customer.FullName,
			PhoneNumber: payload.Customer.PhoneNumber,
			NIK:         payload.Customer.NIK,
			Address:     payload.Customer.Address,
			Status:      payload.Customer.Status,
			BirthDate:   common.FormatDateString(payload.Customer.BirthDate),
			User:        nil,
		},
		ProductType:      payload.ProductType,
		LoanAmount:       payload.LoanAmount,
		Tenure:           payload.Tenure,
		EmploymentStatus: payload.EmploymentStatus,
		MonthlyIncome:    payload.MonthlyIncome,
		Status:           payload.Status,
		RejectionReason:  payload.RejectionReason,
	}
}

func (c *creditUsecase) mapCustomerToResponse(customer model.Customer) dto.CustomerResponseDto {
	return dto.CustomerResponseDto{
		BaseModelResponseDto: dto.BaseModelResponseDto{
			Id: customer.BaseModel.ID,
			CreatedAt: customer.BaseModel.CreatedAt,
			UpdatedAt: customer.BaseModel.UpdatedAt,
		},
		FullName:    customer.FullName,
		PhoneNumber: customer.PhoneNumber,
		NIK:         customer.NIK,
		Address:     customer.Address,
		Status:      customer.Status,
		BirthDate:   common.FormatDateString(customer.BirthDate),
	}
}

func (c *creditUsecase) FindAll(size int, page int) ([]dto.CreditResponseDto, payload.Paging, error) {
	totalRecords, err := c.repo.GetTotal()
	if err != nil {
		return nil, payload.Paging{}, err
	}
	totalPages := (int(totalRecords) + size - 1) / size
	offset := (page - 1) * size
	paging := payload.Paging{
		Page:        page,
		TotalRows:   totalRecords,
		RowsPerPage: size,
		TotalPages:  totalPages,
	}
	credits, err := c.repo.List(size, offset)
	var creditResponses []dto.CreditResponseDto
	for _, credit := range credits {
		creditResponses = append(creditResponses, c.mappingToResponse(credit))
	}
	if err != nil {
		return nil, payload.Paging{}, err
	}
	return creditResponses, paging, nil
}

func NewCreditUsecase(repo repository.CreditRepository) CreditUsecase {
	return &creditUsecase{repo: repo}
}
