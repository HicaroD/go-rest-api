package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllUsersController(ctx echo.Context) error {
	// call service
	return ctx.JSON(http.StatusOK, map[string]string{"Hello from": "GetAllUsersController"})
}
