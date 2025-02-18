package users

import (
	"github.com/labstack/echo/v4"

	"github.com/Viventio/legos/jwt"
	user "lego-api-go/internal/user/services"
)

type Handler struct {
	UserService user.UserService
}

func (h *Handler) RegisterControllers(prefix string, e *echo.Echo) {
	g := e.Group(prefix)
	g.Use(jwt.AuthMiddleware)

	g.GET("/", h.GetAllUsersController)
	g.GET("/:id", h.GetUserByIdController)
	g.POST("/", h.CreateUserController)
}
