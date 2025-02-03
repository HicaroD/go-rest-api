package users

import (
	"github.com/labstack/echo/v4"

	"github.com/Viventio/legos/jwt"
	"lego-api-go/internal/services/users"
)

type Handler struct {
	UserService users.UserService
}

func (h *Handler) RegisterControllers(prefix string, e *echo.Echo) {
	g := e.Group(prefix)
	g.Use(jwt.AuthMiddleware)

	g.GET("/", h.GetAllUsersController)
	g.GET("/:id", h.GetUserByIdController)
	g.POST("/", h.CreateUserController)
}
