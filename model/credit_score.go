package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditScore struct {
	BaseModel
	CustomerID string   `gorm:"type:uuid;not null;unique"` // Relasi ke tabel Customer
	Customer   Customer `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Score           int    `gorm:"not null;check:score BETWEEN 0 AND 100"`                        // Total score customer (0-100)
	Grade           string `gorm:"type:varchar(2);not null;check:grade IN ('A','B','C','D','E')"` // Grade berdasarkan score
	TotalCredits    int    `gorm:"not null"`                                                      // Total jumlah kredit yang dimiliki customer
	ApprovedCredits int    `gorm:"not null"`                                                      // Total kredit yang disetujui
	RejectedCredits int    `gorm:"not null"`                                                      // Total kredit yang ditolak
}

func (cs *CreditScore) BeforeCreate(tx *gorm.DB) (err error) {
	cs.BaseModel.ID = uuid.NewString()
	return
}
