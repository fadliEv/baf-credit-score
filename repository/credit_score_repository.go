package repository

import (
	"baf-credit-score/model"

	"gorm.io/gorm"
)

type CreditScoreRepository interface {
	Save(payload model.CreditScore) error
	Get(id string) (model.CreditScore, error)
	Update(payload model.CreditScore) error
	List(limit int, offset int) ([]model.CreditScore, error)
	GetTotal() (int64, error)
	GetByCustomer(id string) (model.CreditScore, error)
}

type creditScoreRepository struct {
	db *gorm.DB
}

// Get implements CreditScoreRepository.
func (c *creditScoreRepository) Get(id string) (model.CreditScore, error) {
	var creditScore model.CreditScore
	err := c.db.Preload("Customer").Where("id",id).First(&creditScore).Error
	if err != nil {
		return model.CreditScore{},err
	}
	return creditScore,nil
}

// GetByCustomer implements CreditScoreRepository.
func (c *creditScoreRepository) GetByCustomer(id string) (model.CreditScore, error) {
	var creditScore model.CreditScore
	err := c.db.Preload("Customer").Where("customer_id",id).First(&creditScore).Error
	if err != nil {
		return model.CreditScore{},err
	}
	return creditScore,nil
}

// GetTotal implements CreditScoreRepository.
func (c *creditScoreRepository) GetTotal() (int64, error) {
	var count int64
	err := c.db.Model(&model.CreditScore{}).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count,nil
}

// List implements CreditScoreRepository.
func (c *creditScoreRepository) List(limit int, offset int) ([]model.CreditScore, error) {
	var creditScores []model.CreditScore
	if err := c.db.Preload("Customer").Limit(limit).Offset(offset).Find(&creditScores).Error; err != nil {
		return nil,err
	}
	return creditScores,nil
}

// Save implements CreditScoreRepository.
func (c *creditScoreRepository) Save(payload model.CreditScore) error {
	return c.db.Create(&payload).Error
}

// Update implements CreditScoreRepository.
func (c *creditScoreRepository) Update(payload model.CreditScore) error {
	return c.db.Model(&model.CreditScore{}).Where("id = ?",payload.ID).Updates(payload).Error
}

func NewCreditScoreRepository(db *gorm.DB) CreditScoreRepository {
	return &creditScoreRepository{
		db: db,
	}
}
