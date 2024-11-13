package delivery

import (
	"baf-credit-score/config"
	"baf-credit-score/model"
	"baf-credit-score/repository"
	"baf-credit-score/usecase"
	"fmt"
	"time"
)

type server struct {
	customerUsecase usecase.CustomerUsecase
}

func(s *server) Run(){
	// CRUD	
	// Create Customer
	// Siapkan Data Customer 
	newCustomer := model.Customer{
		FullName:    "Ronald",
		PhoneNumber: "098370981184",
		NIK:         "12378123125",
		Address:     "Depok",
		Status:      "active",
		BirthDate:   time.Now(),
	}
	
	err := s.customerUsecase.RegisterCustomer(newCustomer)
	if err != nil {
		fmt.Printf("Error Rgister Customer : %v \n",err)
	}

	// Find All Customer

	// Find Customer By Id

	// Update Customer By Id

	// Delete Customer By Id
}

func NewServer() *server{
	// Ambil Configurasi ENV File untuk kebutuhan Koneksi Database
	cfg,errConfig := config.NewConfig()
	if errConfig != nil {
		fmt.Printf("Error Config : %v \n",errConfig)
	}

	// Init Koneksi dari layer Config > database.go
	db, errConn := config.NewDbConnection(cfg)
	if errConn != nil {
		fmt.Printf("Error Connection : %v \n",errConn)
	}

	fmt.Println("Koneksi aman!!")

	// init customer repository
	repo := repository.NewCustomerRepository(db.Conn())
	// init customer usecase
	usecase := usecase.NewCustomerUsecase(repo)
	return &server{
		customerUsecase: usecase,
	}
}