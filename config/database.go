package config

import (
	"baf-credit-score/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection interface {
	Conn() *gorm.DB
}

type dbConnection struct {
	db *gorm.DB
	cfg *Config
}

func (d *dbConnection) initDbConnection() error {	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	d.cfg.Host,
	d.cfg.User,
	d.cfg.Password,
	d.cfg.Name,
	d.cfg.Port,
	)

	db, errConnection := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errConnection != nil {
		return errConnection
	}
	d.db = db

	// Cek apakah AutoMigration diaktifkan berdasarkan konfigurasi
    if d.cfg.Migration == "true" {
        if err := d.db.AutoMigrate(d.getModels()...); err != nil {
            log.Fatalf("Migration is Failed: %v", err)
        }
        log.Println("Migration is success!")
    }    

	return nil
}

func (d *dbConnection) getModels() []interface{} {
    return []interface{}{
        &model.Customer{},
        &model.User{},    		
		&model.Credit{},
		&model.CreditScore{},
    }
}

func (d *dbConnection) Conn() *gorm.DB{
	return d.db
}

func NewDbConnection(cfg *Config) (DbConnection,error) {
	conn := &dbConnection{
		cfg: cfg,
	}

	err := conn.initDbConnection()
	if err != nil {
		return nil,err
	}
	return conn,nil
}