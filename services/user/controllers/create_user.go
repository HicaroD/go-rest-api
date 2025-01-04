package controllers

import (
	"net/http"

	"github.com/HicaroD/learnanything-api/services/user/repository"
	"github.com/HicaroD/learnanything-api/utils"
	"github.com/labstack/echo/v4"
)

type CreateUserRequestBody struct {
	Name     string `json:"name"       validate:"required"`
	LastName string `json:"last_name"  validate:"required"`
}

func CreateUserController(ctx echo.Context, userRepository repository.User) error {
	createUserRequest := CreateUserRequestBody{}
	err := utils.ValidateRequest(ctx, &createUserRequest)
	if err != nil {
		return err
	}

	// TODO: call some sort of service / repository here

	return ctx.JSON(http.StatusCreated, map[string]string{"Hello from": "CreateUserController"})
}
