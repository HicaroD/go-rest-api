package jwt

import (
	"lego-api-go/legos/cookies"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, parserError := extractJwtToken(c)
		if parserError != nil {
			return parserError
		}

		claims, valid := VerifyToken(token)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		c.Set("user", claims)
		return next(c)
	}
}

func extractJwtToken(c echo.Context) (string, *echo.HTTPError) {
	token, found := cookies.GetCookie(c, "jwt_token")
	if !found {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
	}
	return token.Value, nil
}
