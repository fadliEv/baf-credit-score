package main

import (
	"baf-credit-score/model"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB // Connection untuk melakukan interaksi dengan ORM (migration,crud, dan lain-lain)
var errConnection error

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=baf_credit_score port=5433 sslmode=disable"
	db, errConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errConnection != nil {
		fmt.Printf("Gagal koneksi DB : %v", errConnection)
	}
	fmt.Println("Berhasil Koneksi kedalam Database")
	// Auto Migration
	// err := db.AutoMigrate(
	// 	&model.Customer{},
	// 	&model.User{},
	// )

	// err = db.Migrator().DropColumn(&model.Customer{},"Status") // Untuk menghapus column yang sudah ada
	// if err != nil {
	// 	fmt.Printf("Gagal Hapus Column : %v",err)
	// }
	// if err != nil {
	// 	fmt.Printf("Gagal Migration : %v",err)
	// }
	// registerCustomer()
	// getAllCustomer()

	// var customers model.Customer
	// if err := db.Table("customers").First(&customers); err != nil {
	// 	fmt.Printf("failed get All customer : %v", err.Error.Error())
	// 	return
	// }
	// fmt.Println(customers)
	// getAllCustomer()
	findCustomerById()
}

func registerCustomer() { // Fuction untuk insert data customer kedalam database
	// sql.exec("inser into...",value,value)
	birthDateStr := "01-11-2001"
	format := "02-01-2006"
	birthDate, err := time.Parse(format,birthDateStr)
	if err != nil {
		fmt.Printf("failed parse birthDate : %v", err)
		return
	}
	newCustomer := model.Customer{
		FullName:    "Juan",
		PhoneNumber: "098378921184",
		NIK:         "12378901125",
		Address:     "Depok",
		Status:      "active",
		BirthDate:   birthDate,
	}
	errCreate := db.Create(&newCustomer)
	if errCreate.Error != nil {
		fmt.Printf("failed insert customer : %v", errCreate.Error)
		return
	}
	fmt.Println("Success register customer")
}

func getAllCustomer() { // ambil semua record dari table
	var customers []model.Customer
	err := db.Find(&customers); 
	if err.Error != nil {
		fmt.Printf("failed get All customer : %v", err.Error.Error())
		return
	}
	for _, customer := range customers {
		fmt.Println(customer)
	}

	// Menghapus column tertentu
	// err = db.Migrator().DropColumn(&model.Customer{},"Status") // Untuk menghapus column yang sudah ada
	// if err != nil {
	// 	fmt.Printf("Gagal Hapus Column : %v",err)
	// }
}

func findCustomerById(){
	var customer model.Customer
	err := db.First(&customer,"id = ?","5fe6f6bb-a82d-42fc-9aa4-4ed3105c2be9"); 
	if err.Error != nil {
		fmt.Printf("failed get customer by id : %v", err.Error.Error())
		return
	}
	fmt.Println("Customer :", customer)
}