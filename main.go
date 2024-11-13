package main

import (	
	"baf-credit-score/delivery"
)

func main() {
	delivery.NewServer().Run()
}

// func registerCustomer() { // Fuction untuk insert data customer kedalam database
// 	// sql.exec("inser into...",value,value)
// 	birthDateStr := "01-11-2001"
// 	format := "02-01-2006"
// 	birthDate, err := time.Parse(format,birthDateStr)
// 	if err != nil {
// 		fmt.Printf("failed parse birthDate : %v", err)
// 		return
// 	}
// 	newCustomer := model.Customer{
// 		FullName:    "Juan",
// 		PhoneNumber: "098378921184",
// 		NIK:         "12378901125",
// 		Address:     "Depok",
// 		Status:      "active",
// 		BirthDate:   birthDate,
// 	}
// 	errCreate := db.Create(&newCustomer)
// 	if errCreate.Error != nil {
// 		fmt.Printf("failed insert customer : %v", errCreate.Error)
// 		return
// 	}
// 	fmt.Println("Success register customer")
// }

// func getAllCustomer() { // ambil semua record dari table
// 	var customers []model.Customer
// 	err := db.Find(&customers); 
// 	if err.Error != nil {
// 		fmt.Printf("failed get All customer : %v", err.Error.Error())
// 		return
// 	}
// 	for _, customer := range customers {
// 		fmt.Println(customer)
// 	}
// }

// func findCustomerById(){
// 	var customer model.Customer
// 	err := db.First(&customer,"id = ?","fc973770-fd88-4877-ada8-353687e000b7"); 
// 	if err.Error != nil {
// 		fmt.Printf("failed get customer by id : %v", err.Error.Error())
// 		return
// 	}
// 	fmt.Println("Customer :", customer)
// }