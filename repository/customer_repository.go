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
	List(limit int, offset int) ([]model.Customer, error)       //Get semua data
	Delete(id string) error                //Hapus data berdasarkan ID
	GetTotal() (int64, error) 			   //Total records
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

func (c *customerRepository) List(limit int, offset int) ([]model.Customer, error) {
	var customers []model.Customer
    if err := c.db.Preload("User").Limit(limit).Offset(offset).Find(&customers).Error; err != nil {        
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

func (c *customerRepository) GetTotal() (int64,error) {	
	var count int64
    if err := c.db.Model(&model.Customer{}).Count(&count).Error; err != nil {        
        return 0, err
    }
    return count, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
