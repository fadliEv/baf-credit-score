package manager

import (
	"baf-credit-score/config"
	"baf-credit-score/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
	Migrate(model ...any) error
	Config() *config.Config
}

type infraManager struct {
	db *gorm.DB
	cfg *config.Config
}

func (i *infraManager) Conn() *gorm.DB {
	return i.db
}

func (i *infraManager)  Config() *config.Config {
	return i.cfg
}

func (i *infraManager) Migrate(model ...any) error {
	err := i.db.AutoMigrate(model...)
	if err != nil {
		return err
	}
	return nil
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	i.db = conn
	if i.cfg.DbConfig.Migration == "MIGRATION" {
		i.db = conn.Debug()
		err := i.Migrate(i.getModels()...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *infraManager) getModels() []interface{} {
    return []interface{}{
        &model.Customer{},
        &model.User{},    		
		&model.Credit{},
		&model.CreditScore{},
    }
}


func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}