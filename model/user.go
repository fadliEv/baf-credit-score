package model

type User struct {
	BaseModel
	Email    string `gorm:"type:varchar(50);not null;uniqueIndex"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:varchar(20)"`
}
