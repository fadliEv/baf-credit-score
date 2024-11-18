package usecase

import (
	"baf-credit-score/model"
	"baf-credit-score/model/dto"
	"baf-credit-score/repository"	
	"baf-credit-score/utils/common"
	"baf-credit-score/utils/constant"
	"baf-credit-score/utils/payload"
)

type CustomerUsecase interface {
	RegisterCustomer(payload dto.CustomerRequestDto) (dto.CustomerResponseDto,error)
	FindCustomerById(id string) (dto.CustomerResponseDto, error)
	FindAll(size int, page int) ([]dto.CustomerResponseDto,payload.Paging,error) 
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
func (c *customerUsecase) FindAll(size int, page int) ([]dto.CustomerResponseDto,payload.Paging,error) {
	totalRecords, err := c.repo.GetTotal()
	if err != nil {
		return nil, payload.Paging{},err
	}
	totalPages := (int(totalRecords) + size - 1) / size
	offset := (page - 1) * size
	paging := payload.Paging {
		Page: page,
		TotalRows: totalRecords,
		RowsPerPage: size,
		TotalPages: totalPages,
	}
	customers, err := c.repo.List(size,offset)
	var customerResponses []dto.CustomerResponseDto
	for _, customer := range customers {		
		customerResponses =	append(customerResponses,c.mappingResponse(customer))
	}
	if err != nil {
		return nil, payload.Paging{},err
	}
	return customerResponses,paging,nil
}

// UpdateCustomer implements CustomerUsecase.
func (c *customerUsecase) UpdateCustomer(payload dto.CustomerRequestDto) error {
	customerModel, err := c.mappingRequest(payload)
	if err != nil {
		return err
	}
	customerModel.ID = payload.Id
	return c.repo.Update(customerModel)
}

// FindCustomerById implements CustomerUsecase.
func (c *customerUsecase) FindCustomerById(id string) (dto.CustomerResponseDto, error) {
	customer, err := c.repo.Get(id)
	if err != nil {
		return dto.CustomerResponseDto{}, err
	}
	return c.mappingResponse(customer),nil
}

// RegisterCustomer implements CustomerUsecase.
func (c *customerUsecase) RegisterCustomer(payload dto.CustomerRequestDto) (dto.CustomerResponseDto,error){
	customerModel, err := c.mappingRequest(payload)
	if err != nil {
		return dto.CustomerResponseDto{},err
	}
	hashedPass, err := common.HashPassword(payload.Password)
	if err != nil {
		return dto.CustomerResponseDto{},err
	}
	user := model.User{
		Email:    payload.Email,
		Password: hashedPass,
		Role:     constant.USER,
	}
	customerModel.User = user	
	errSave := c.repo.Save(customerModel)
	if errSave != nil {
		return dto.CustomerResponseDto{},errSave
	}
	customerResponse := dto.CustomerResponseDto{
		FullName: payload.FullName,
		PhoneNumber: payload.PhoneNumber,
		NIK: payload.NIK,
		Address: payload.Address,
		Status: payload.Status,
		BirthDate: payload.BirthDate,
		User: dto.UserResponseDto{
			Email: payload.Email,
			Role: constant.USER,
		},
	}
	return customerResponse,nil
}

func (c *customerUsecase) mappingRequest(payload dto.CustomerRequestDto) (model.Customer,error){
	parseBirthDate, err := common.ParseDate(payload.BirthDate)
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

func (c *customerUsecase) mappingResponse(payload model.Customer) dto.CustomerResponseDto {
	return dto.CustomerResponseDto{
		BaseModelResponseDto: dto.BaseModelResponseDto{
			Id: payload.ID,
			CreatedAt: payload.CreatedAt,
			UpdatedAt: payload.UpdatedAt,
		},
		FullName: payload.FullName,
		PhoneNumber: payload.PhoneNumber,
		NIK: payload.NIK,
		Address: payload.Address,
		Status: payload.Status,
		BirthDate: common.FormatDateString(payload.BirthDate),
		User: dto.UserResponseDto{
			BaseModelResponseDto: dto.BaseModelResponseDto{
				Id: payload.User.ID,
				CreatedAt: payload.User.CreatedAt,
				UpdatedAt: payload.User.UpdatedAt,
			},
			Email: payload.User.Email,
			Role: payload.User.Role,
		},
	}
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		repo: repo,
	}
}
