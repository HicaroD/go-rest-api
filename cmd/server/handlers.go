package main

import (
	userH "github.com/HicaroD/api/internal/user/controllers"
	"github.com/HicaroD/api/internal/user/service"
	"github.com/labstack/echo/v4"
)

func registerAllHandlers(e *echo.Echo) {
	// TODO: initialize databases, loggers, configuration files and more

	userRepository := service.NewUserService()
	userHandler := userH.NewHandler(userRepository)
	userHandler.RegisterControllers("/users", e)
}
