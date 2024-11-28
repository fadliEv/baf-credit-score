package delivery

import (
	"baf-credit-score/config"
	"baf-credit-score/delivery/controller"
	"baf-credit-score/delivery/middleware"
	"baf-credit-score/repository"
	"baf-credit-score/usecase"
	"baf-credit-score/utils/service"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

type server struct {
	customerUsecase usecase.CustomerUsecase
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthenticationUsecase
	creditUsecase usecase.CreditUsecase
	creditScoreUsecase usecase.CreditScoreUsecase
	engine *gin.Engine
	jwtSerivce service.JwtService
}

func(s *server) setupController(){
	authMiddleware := middleware.NewAuthMiddleware(s.jwtSerivce)
	rg := s.engine.Group("/api/v1")	
	controller.NewCustomerController(s.customerUsecase,rg,authMiddleware).Route()
	controller.NewUserController(s.userUsecase,rg,authMiddleware).Route()
	controller.NewAuthController(s.authUsecase,rg).Route()
	controller.NewCreditController(s.creditUsecase,rg,authMiddleware).Route()
	controller.NewCreditScoreController(s.creditScoreUsecase,rg,authMiddleware).Route()
}

func(s *server) Run(){
	s.setupController()
	if err := s.engine.Run(); err != nil {
		panic(fmt.Errorf("server not running %s ",err.Error()))
	}
}

func NewServer() *server{
	// intance gin engine
	f, err := os.Create("api.log")
	if err != nil {
		panic(errors.New("failed init log file"))
	}
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
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

	// -------------------------------------- Customer	
	customerRepo := repository.NewCustomerRepository(db.Conn())	
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)


	// -------------------------------------- User	
	userRepo := repository.NewUserRepository(db.Conn())
	userUsecase := usecase.NewUserUsecase(userRepo)	

	// test jwt service
	jwtService := service.NewJwtService(cfg.TokenConfig)

	// -------------------------------------- Auth
	authUsecase := usecase.NewAuthenticationUsecase(userUsecase,jwtService)	

	
	// -------------------------------------- Credit Score
	creditScoreRepo := repository.NewCreditScoreRepository(db.Conn())
	creditScoreUsecase := usecase.NewCreditScoreUsecase(creditScoreRepo)
	
	// -------------------------------------- Credit
	creditRepo := repository.NewCreditRepository(db.Conn())
	creditUsecase := usecase.NewCreditUsecase(creditRepo,creditScoreUsecase)

	return &server{
		customerUsecase: customerUsecase,
		userUsecase: userUsecase,
		engine: ginEngine, // assign ke dalam struct server
		jwtSerivce: jwtService,
		authUsecase: authUsecase,
		creditUsecase: creditUsecase,
		creditScoreUsecase: creditScoreUsecase,
	}
}