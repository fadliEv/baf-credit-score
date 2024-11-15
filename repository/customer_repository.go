package repository

import (
	"baf-credit-score/model"
	"gorm.io/gorm"
)

/**
INTERFACE
____________

STRUCT

func
func
...
____________

CONSTRUCTOR

**/

type CustomerRepository interface {
	InsertCustomer(newCustomer model.Customer) error
	FindCustomerById(id string) (model.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

// FindCustomerById implements CustomerRepository.
func (c *customerRepository) FindCustomerById(id string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.First(&customer,"id = ?",id).Error
	if err != nil {
		return model.Customer{},err
	}
	return customer,nil
}

// insertCustomer implements CustomerRepository.
func (c *customerRepository) InsertCustomer(newCustomer model.Customer) error {
	// insert data kedalam database
	// sql.exec("inser into...",value,value)
	errCreate := c.db.Create(&newCustomer)
	if errCreate.Error != nil {
		return errCreate.Error
	}
	return nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
