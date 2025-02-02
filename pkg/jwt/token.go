package jwt

import (
	"lego-api-go/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload map[string]string, expirationTimeInSeconds int64) (string, error) {
	claims := jwt.MapClaims{
		"exp": expirationTimeInSeconds,
	}
	for k, v := range payload {
		claims[k] = v
	}
	// TODO: set signing method properly
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.CFG.JWTSecret))
}

func VerifyToken(token string) (jwt.MapClaims, bool) {
	claims := jwt.MapClaims{}
	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.CFG.JWTSecret), nil
	})
	if err != nil {
		return nil, false
	}
	if !tk.Valid {
		return nil, false
	}

	return claims, true
}
