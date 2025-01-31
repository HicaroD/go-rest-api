package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload map[string]any, expirationTimeInSeconds int64) (string, error) {
	claims := jwt.MapClaims{
		"exp": expirationTimeInSeconds,
	}
	for k, v := range payload {
		claims[k] = v
	}
	// TODO: set signing method properly
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
