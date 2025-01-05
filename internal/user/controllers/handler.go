package controllers

import (
	"github.com/labstack/echo/v4"

	"github.com/HicaroD/learnanything-api/internal/user/service"
)

type Handler struct {
	userService service.User
}

func NewHandler(userService service.User) *Handler {
	return &Handler{userService}
}

func (h *Handler) RegisterControllers(prefix string, e *echo.Echo) {
	g := e.Group(prefix)

	g.GET("/", h.GetAllUsersController)
	g.POST("/:id", h.GetUserByIdController)
	g.POST("/",h.CreateUserController)
}
