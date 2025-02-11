package login

import (
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (h *Handler) RegisterControllers(prefix string, e *echo.Echo) {
	g := e.Group(prefix)
	g.POST("/", h.LoginController)
}
