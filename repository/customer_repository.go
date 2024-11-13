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
}

type customerRepository struct{
	db *gorm.DB
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