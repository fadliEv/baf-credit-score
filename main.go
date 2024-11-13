package main

import (
	"baf-credit-score/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=baf_credit_score port=5433 sslmode=disable"
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Printf("Gagal koneksi DB : %v",err)
	}
	fmt.Println("Berhasil Koneksi kedalam Database")

	// Auto Migration
	err = db.AutoMigrate(
		&model.Customer{}, 
		&model.User{},
	)

	// err = db.Migrator().DropColumn(&model.Customer{},"Status") // Untuk menghapus column yang sudah ada
	if err != nil {
		fmt.Printf("Gagal Hapus Column : %v",err)
	}
	if err != nil {
		fmt.Printf("Gagal Migration : %v",err)
	}
}


/**
01-connection
02-declare-modes
03-clean-arch
..
unit-testing
jwt
**/