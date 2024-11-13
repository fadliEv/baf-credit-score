package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=baf_credit_score port=5433 sslmode=disable"
	_,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Printf("Gagal koneksi DB : %v",err)
	}
	fmt.Println("Berhasil Koneksi kedalam Database")
}


/**
01-connection
02-declare-modes
03-clean-arch
..
unit-testing
jwt
**/