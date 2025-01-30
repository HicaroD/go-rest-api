package users

import (
	"github.com/labstack/echo/v4"

	"lego-api-go/internal/services/users"
)

type Handler struct {
	UserService users.UserService
}

func (h *Handler) RegisterControllers(prefix string, e *echo.Echo) {
	g := e.Group(prefix)

	g.GET("/", h.GetAllUsersController)
	g.GET("/:id", h.GetUserByIdController)
	g.POST("/", h.CreateUserController)
}
