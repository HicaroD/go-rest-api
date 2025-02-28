package login

import (
	"github.com/Viventio/legos/cookies"
	"github.com/Viventio/legos/crypt"
	"github.com/Viventio/legos/jwt"
	"github.com/Viventio/legos/validators"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LoginRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var oneWeek int64 = time.Now().AddDate(0, 0, 7).Unix()

func (h *Handler) LoginController(ctx echo.Context) error {
	req := LoginRequestBody{}
	err := validators.ValidateRequest(ctx, &req)
	if err != nil {
		return err
	}

	// TODO: read hashed user password from db
	hashedPasswordFromDB, _ := crypt.HashPassword(req.Password)
	if !crypt.ComparePassword(req.Password, hashedPasswordFromDB) {
		return ctx.JSON(http.StatusUnauthorized, map[string]any{
			"detail": "invalid credentials",
		})
	}

	// TODO: read user data for db and set here
	payload := map[string]any{
		"user_id":  10,
		"is_admin": false,
	}

	token, err := jwt.GenerateToken(payload, oneWeek)
	if err != nil {
		return err
	}
	cookies.SetCookie(ctx, "jwt_token", token, true, oneWeek)

	return ctx.JSON(http.StatusOK, map[string]any{
		"detail": "User logged in successfuly",
	})
}
