package controllers

import (
	"net/http"

	"github.com/HicaroD/learnanything-api/utils"
	"github.com/labstack/echo/v4"
)

type GetUserByIdRequestBody struct {
	ID   string `validate:"required"`
	Name string `json:"name" validate:"required"`
}

func GetUserByIdController(ctx echo.Context) error {
	// TODO: it would be awesome to extract all necessary
	// informations from the request (param, query, body and more)
	// in one single command

	req := &GetUserByIdRequestBody{
		ID: ctx.Param("id"),
	}
	err := utils.ValidateRequest(ctx, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"Hello from": "GetUserByIdController - id: " + req.ID})
}
