package repository

import (
	"baf-credit-score/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(payload model.User) error
	Get(id string) (model.User, error)
	GetByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// Get implements UserRepository.
func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.Where("id=?",id).First(&user).Error
	if err != nil {
		return model.User{},err
	}
	return user,nil
}

// GetByEmail implements UserRepository.
func (u *userRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Where("email=?",email).First(&user).Error
	if err != nil {
		return model.User{},err
	}
	return user,nil
}

// Save implements UserRepository.
func (u *userRepository) Save(payload model.User) error {
	return u.db.Create(&payload).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
