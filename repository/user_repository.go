package repository

import (
	"baf-credit-score/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(payload model.User) error
	Get(id string) (model.User, error)
	GetByEmail(email string) (model.User, error)
	List(limit int, offset int) ([]model.User, error)
	GetTotal() (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

// Get implements UserRepository.
func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetByEmail implements UserRepository.
func (u *userRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// Save implements UserRepository.
func (u *userRepository) Save(payload model.User) error {
	return u.db.Create(&payload).Error
}

func (u *userRepository) List(limit int, offset int) ([]model.User, error) {
	var users []model.User
	if err := u.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (c *userRepository) GetTotal() (int64, error) {
	var count int64
	if err := c.db.Model(&model.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
