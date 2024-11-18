package service

import (
	"baf-credit-score/config"
	"baf-credit-score/model"
	"baf-credit-score/utils/payload"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	CreateAccessToken(credential model.User) (string, error) // Membuat token ketika client request login
	VerifyAccessToken(tokenString string) (jwt.MapClaims,error)                // Verifikasi token apakah sesuai dengan datanya (signatureKey,mapClaims)
}

type jwtService struct{
	cfg config.TokenConfig
}

// Membuat Token
func (j *jwtService) CreateAccessToken(credential model.User) (string, error) {
	// Waktu ketika token ini dibuat
	now := time.Now().UTC()
	// add time, untuk expired tokennya
	end := now.Add(j.cfg.AccessTokenLifeTime)
	// standard object untuk payload di dalam token
	claims := payload.MyClaim {
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: j.cfg.ApplicationName,
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Id : credential.ID,
		Role: credential.Role,
	}
	jwtNewClaim := jwt.NewWithClaims(j.cfg.JwtSigningMethod,claims)
	token, err := jwtNewClaim.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return "",err
	}
	return token,nil
}

// verfikasi token
func (j *jwtService) VerifyAccessToken(tokenString string) (jwt.MapClaims,error) {
	token, err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != j.cfg.JwtSigningMethod {
			return nil, errors.New("signing method invalid")
		}
		return j.cfg.JwtSignatureKey,nil
	})
	if err != nil {
		return nil,errors.New("Failed parse token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != j.cfg.ApplicationName {
		return nil,errors.New("invalid token claims")
	}
	return claims,nil
}

func NewJwtService(configArgs config.TokenConfig) JwtService {
	return &jwtService{
		cfg:configArgs,
	}
}
