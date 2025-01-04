package main

import (
	"github.com/HicaroD/learnanything-api/services/user"
	"github.com/HicaroD/learnanything-api/services/user/repository"
	"github.com/labstack/echo/v4"
)

func registerAllHandlers(e *echo.Echo) {
	userRepository := repository.NewUserRepository()
	userHandler := user.NewHandler(userRepository)
	userHandler.RegisterRoutes("/users", e)
}
