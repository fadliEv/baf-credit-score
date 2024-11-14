package model

import (
	"time"
)

type Customer struct {
	BaseModel
	FullName    string    `gorm:"type:varchar(50);not null"`
	PhoneNumber string    `gorm:"type:varchar(15);not null;uniqueIndex"`
	NIK         string    `gorm:"type:varchar(16);not null;uniqueIndex"`
	Address     string    `gorm:"type:text"`
	Status      string    `gorm:"type:varchar(10);not null;check:status IN ('active','inactive')"`
	BirthDate   time.Time `gorm:"type:date;not null"`
}
