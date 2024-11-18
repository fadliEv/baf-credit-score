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
	Save(payload model.Customer) error     //Insert data
	Get(id string) (model.Customer, error) //Get data berdasarkan ID
	Update(payload model.Customer) error   //Update data
	List() ([]model.Customer, error)       //Get semua data
	Delete(id string) error                //Hapus data berdasarkan ID
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Get(id string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.Where("id", id).First(&customer).Error
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) List() ([]model.Customer, error) {
	var customers []model.Customer
	err := c.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerRepository) Save(payload model.Customer) error {
	return c.db.Create(&payload).Error
}

func (c *customerRepository) Update(payload model.Customer) error {
	return c.db.Model(&model.Customer{}).Where("id = ?", payload.ID).Updates(payload).Error
}

func (c *customerRepository) Delete(id string) error {	
	return c.db.Delete(&model.Customer{}, "id = ?", id).Error
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
