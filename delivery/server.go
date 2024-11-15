package delivery

import (
	"baf-credit-score/config"
	"baf-credit-score/delivery/controller"
	"baf-credit-score/repository"
	"baf-credit-score/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type server struct {
	customerUsecase usecase.CustomerUsecase
	engine *gin.Engine
}

func(s *server) setupController(){
	rg := s.engine.Group("/api/v1")
	controller.NewCustomerController(s.customerUsecase,rg).Route()
}

func(s *server) Run(){
	s.setupController()
	if err := s.engine.Run(); err != nil {
		panic(fmt.Errorf("server not running %s ",err.Error()))
	}
}

func NewServer() *server{
	// intance gin engine
	ginEngine := gin.Default()

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

	// init customer repository
	repo := repository.NewCustomerRepository(db.Conn())
	
	// init customer usecase
	usecase := usecase.NewCustomerUsecase(repo)
	return &server{
		customerUsecase: usecase,
		engine: ginEngine, // assign ke dalam struct server
	}
}