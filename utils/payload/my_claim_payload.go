package payload

import "github.com/golang-jwt/jwt/v5"

type MyClaim struct {
	jwt.RegisteredClaims
	// role
	Id   string `json:"id"`
	Role string `json:"role"`
}
