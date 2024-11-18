package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host      string
	Port      string
	Name      string
	User      string
	Password  string
	Driver    string
	Migration string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	DbConfig
	TokenConfig
}

func (c *Config) ReadConfigFile() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error load Env File : %v", err.Error())
		return err
	}
	c.DbConfig = DbConfig{
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Name:      os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		Driver:    os.Getenv("DB_DRIVER"),
		Migration: os.Getenv("MIGRATION"),
	}

	// App Name
	// Expired Time
	appTokenExp, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRED"))
	if err != nil {
		return err
	}
	accessTokenLifeTime := time.Duration(appTokenExp) * time.Minute
	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("APP_TOKEN_NAME"),
		JwtSignatureKey:     []byte(os.Getenv("APP_TOKEN_KEY")),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.ReadConfigFile(); err != nil {
		return nil, err
	}
	return cfg, nil
}
