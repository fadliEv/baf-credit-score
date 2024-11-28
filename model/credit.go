package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Credit struct {
	BaseModel
	CustomerID       string   `gorm:"not null"`
	Customer         Customer `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AppNumber        string   `gorm:"type:varchar(50);not null;uniqueIndex"`                                       // Nomor aplikasi kredit
	ProductType      string   `gorm:"type:varchar(50);not null;check:product_type IN ('CAR','MOTOR','FURNITURE')"`       // Jenis produk kredit
	LoanAmount       float64  `gorm:"not null"`                                                                    // Jumlah pinjaman
	Tenure           int      `gorm:"not null"`                                                                    // Lama tenor dalam bulan
	EmploymentStatus string   `gorm:"type:varchar(20);not null;check:employment_status IN ('PKWT','PKWTT')"`                  // Status pekerjaan customer
	MonthlyIncome    float64  `gorm:"not null"`                                                                    // Pendapatan bulanan
	Status           string   `gorm:"type:varchar(20);not null;check:status IN ('PENDING','APPROVED','REJECTED')"` // Status pengajuan
	RejectionReason  string   `gorm:"type:varchar(255)"`                                                           // Alasan penolakan (jika ada)
}

func (c *Credit) BeforeCreate(tx *gorm.DB) (err error) {
	c.BaseModel.ID = uuid.NewString()
	return
}
