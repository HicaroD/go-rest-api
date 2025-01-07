package controllers

import (
	"net/http"

	"github.com/HicaroD/api/pkg/utils"
	"github.com/labstack/echo/v4"
)

type CreateUserRequestBody struct {
	Name     string `json:"name"       validate:"required"`
	LastName string `json:"last_name"  validate:"required"`
}

func (h *Handler) CreateUserController(ctx echo.Context) error {
	createUserRequest := CreateUserRequestBody{}
	err := utils.ValidateRequest(ctx, &createUserRequest)
	if err != nil {
		return err
	}

	// TODO: call some sort of service here

	return ctx.JSON(http.StatusCreated, map[string]string{"Hello from": "CreateUserController"})
}
