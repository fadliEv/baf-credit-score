package delivery

import (
	"baf-credit-score/config"
	"baf-credit-score/delivery/controller"
	"baf-credit-score/delivery/middleware"
	"baf-credit-score/manager"
	"baf-credit-score/utils/service"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type server struct {
	ucManager  manager.UsecaseManager
	engine     *gin.Engine
	host       string
	jwtSerivce service.JwtService
}

func (s *server) setupController() {
	authMiddleware := middleware.NewAuthMiddleware(s.jwtSerivce)
	rg := s.engine.Group("/api/v1")
	controller.NewCustomerController(s.ucManager.CustomerUsecase(), rg, authMiddleware).Route()
	controller.NewUserController(s.ucManager.UserUsecas(), rg, authMiddleware).Route()
	controller.NewAuthController(s.ucManager.AuthUsecase(), rg).Route()
	controller.NewCreditController(s.ucManager.CreditUsecase(), rg, authMiddleware).Route()
	controller.NewCreditScoreController(s.ucManager.CreditScoreUsecase(), rg, authMiddleware).Route()
}

func (s *server) Run() {
	s.setupController()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running %s ", err.Error()))
	}
}

func NewServer() *server {
	// intance gin engine
	f, err := os.Create("api.log")
	if err != nil {
		log.Printf("failed init log file :%s", err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	ginEngine := gin.Default()

	// Ambil Configurasi ENV File untuk kebutuhan Koneksi Database
	cfg, errConfig := config.NewConfig()
	if errConfig != nil {
		fmt.Printf("Error Config : %v \n", errConfig)
	}

	// init infra manager
	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Printf("Error init Infra : %v \n", err)
	}

	// init repo manager
	repo := manager.NewRepositoryManager(infra)

	// init usecase manager
	usecase := manager.NewUsecaseManager(repo)

	// Host
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)

	// Jwt Service
	jwtService := service.NewJwtService(cfg.TokenConfig)

	return &server{
		ucManager: usecase,
		engine:    ginEngine, // assign ke dalam struct server
		host:      host,
		jwtSerivce: jwtService,
	}
}
