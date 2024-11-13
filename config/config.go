package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type Config struct {
	DbConfig
}

func (c *Config) ReadConfigFile() error {
	err := godotenv.Load()
	if err != nil{
		fmt.Printf("Error load Env File : %v",err.Error())
		return err
	}
	c.DbConfig = DbConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver: os.Getenv("DB_DRIVER"),
	}
	return nil
}

func NewConfig()(*Config,error){
	cfg := &Config{}
	if err := cfg.ReadConfigFile(); err != nil {
		return nil,err
	}
	return cfg,nil
}