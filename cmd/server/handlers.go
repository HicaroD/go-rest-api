package main

import (
	"github.com/HicaroD/learnanything-api/services/user"
	"github.com/HicaroD/learnanything-api/services/user/repository"
	"github.com/labstack/echo/v4"
)

func registerAllHandlers(e *echo.Echo) {
	// TODO: initialize databases, loggers, configuration files and more

	userRepository := repository.NewUserRepository()
	userHandler := user.NewHandler(userRepository)
	userHandler.RegisterRoutes("/users", e)
}
