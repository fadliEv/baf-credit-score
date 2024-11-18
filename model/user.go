package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Email    string `gorm:"type:varchar(50);not null;uniqueIndex"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:varchar(20)"`
}

func (c *User) BeforeCreate(tx *gorm.DB) (err error) {
	c.BaseModel.ID = uuid.NewString()
	return
}