package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	BaseModel
	FullName    string    `gorm:"type:varchar(50);not null"`
	PhoneNumber string    `gorm:"type:varchar(15);not null;uniqueIndex"`
	NIK         string    `gorm:"type:varchar(16);not null;uniqueIndex"`
	Address     string    `gorm:"type:text"`
	Status      string    `gorm:"type:varchar(10);not null;check:status IN ('active','inactive')"`
	BirthDate   time.Time `gorm:"type:date;not null"`
	UserID      string    `gorm:"not null;unique"`
    User        User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB)(err error){
	c.BaseModel.ID = uuid.NewString()
	return
}