package login

import (
	"lego-api-go/pkg/cookies"
	"lego-api-go/pkg/jwt"
	"lego-api-go/pkg/validators"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LoginRequestBody struct {
	username string `param:"username" validate:"required"`
	password string `param:"password" validate:"required"`
}

var oneWeek int64 = time.Now().AddDate(0, 0, 7).Unix()

func (h *Handler) LoginController(ctx echo.Context) error {
	req := LoginRequestBody{}
	err := validators.ValidateRequest(ctx, &req)
	if err != nil {
		return err
	}

	payload := map[string]any{
		"username": req.username,
		"password": req.password,
	}

	token, err := jwt.GenerateToken(payload, oneWeek)
	if err != nil {
		return err
	}

	cookies.SetSecureCookie(ctx, "jwt_token", token, oneWeek)

	return ctx.JSON(http.StatusOK, map[string]any{
		"detail": "User logged successfuly",
		"token":  token,
	})
}
