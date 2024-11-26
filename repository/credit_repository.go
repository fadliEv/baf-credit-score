package repository

import (
	"baf-credit-score/model"
	"gorm.io/gorm"
)

type CreditRepository interface {
	Save(payload model.Credit) error
	Get(id string) (model.Credit, error)
	Update(payload model.Credit) error
	List(limit int, offset int) ([]model.Credit, error)
	GetTotal() (int64, error)
	GetByCustomer(id string) ([]model.Credit,error) 
}

type creditRepository struct {
	db *gorm.DB
}

func (c *creditRepository) Get(id string) (model.Credit, error) {
	var credit model.Credit
	err := c.db.Where("id", id).First(&credit).Error
	if err != nil {
		return model.Credit{}, err
	}
	return credit, nil
}

func (c *creditRepository) List(limit int, offset int) ([]model.Credit, error) {
	var credits []model.Credit
	if err := c.db.Preload("Customer").Limit(limit).Offset(offset).Find(&credits).Error; err != nil {
		return nil, err
	}
	return credits, nil
}	

func (c *creditRepository) Save(payload model.Credit) error {
	return c.db.Create(&payload).Error
}

func (c *creditRepository) Update(payload model.Credit) error {
	return c.db.Model(&model.Credit{}).Where("id = ?", payload.ID).Updates(payload).Error
}

func (c *creditRepository) GetTotal() (int64, error) {
	var count int64
	if err := c.db.Model(&model.Credit{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (c *creditRepository) GetByCustomer(id string) ([]model.Credit,error) {
	var credits []model.Credit
	err := c.db.Preload("Customer").Where("customer_id = ?", id).Find(&credits).Error
	if err != nil {
		return nil,err
	}
	return credits,nil
}

func NewCreditRepository(db *gorm.DB) CreditRepository {
	return &creditRepository{db: db}
}
