package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllUsersController(ctx echo.Context) error {
	// TODO: call the service here
	return ctx.JSON(http.StatusOK, map[string]string{"Hello from": "GetAllUsersController"})
}
