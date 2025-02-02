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

	// TODO: find username in the database
	// TODO: compare hashed password to the passed password
	// TODO: if comparison is true, generate jwt token with user id (only)
	// TODO: if comparison is false, then unauthorized

	payload := map[string]string{
		"username": req.Username,
		"password": req.Password,
	}
	token, err := jwt.GenerateToken(payload, oneWeek)
	if err != nil {
		return err
	}
	cookies.SetSecureCookie(ctx, "jwt_token", token, oneWeek)

	return ctx.JSON(http.StatusOK, map[string]any{
		"detail": "User logged in successfuly",
	})
}
