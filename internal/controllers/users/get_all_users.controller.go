package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllUsersController(ctx echo.Context) error {
	// TODO: call some sort of service here

	// TODO: build a package around it for later usage
	// cookie := new(http.Cookie)
	// cookie.Name = "user_id"
	// cookie.Value = "102012"
	// cookie.Secure = true
	// cookie.HttpOnly = true
	// ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, map[string]string{"Hello from": "GetAllUsersController"})
}
