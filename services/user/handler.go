package user

import (
	"github.com/labstack/echo/v4"

	"github.com/HicaroD/learnanything-api/services/user/controllers"
	"github.com/HicaroD/learnanything-api/services/user/repository"
)

type Handler struct {
	userRepository repository.User
}

func NewHandler(userRepository repository.User) *Handler {
	return &Handler{userRepository}
}

func (h *Handler) RegisterRoutes(prefix string, e *echo.Echo) {
	g := e.Group(prefix)

	// NOTE: I'm not calling the controller directly because I
	// might use pass some parameters for dependency injection
	g.GET("/", func(ctx echo.Context) error {
		return controllers.GetAllUsersController(ctx)
	})

	g.POST("/:id", func(ctx echo.Context) error {
		return controllers.GetUserByIdController(ctx)
	})

	g.POST("/", func(ctx echo.Context) error {
		return controllers.CreateUserController(ctx, h.userRepository)
	})
}
