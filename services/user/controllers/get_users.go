package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsersController(ctx echo.Context) error {
	// TODO: call some sort of service / repository here
	return ctx.JSON(http.StatusOK, map[string]string{"Hello from": "GetAllUsersController"})
}
