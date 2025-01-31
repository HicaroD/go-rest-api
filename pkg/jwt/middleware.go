package jwt

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// TODO: read from env file
const jwtSecret = "SECRET_KEY"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString, parserError := extractJwtToken(c)
		if parserError != nil {
			return parserError
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user", token)

		return next(c)
	}
}

func extractJwtToken(c echo.Context) (string, *echo.HTTPError) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
	}

	return parts[1], nil
}
