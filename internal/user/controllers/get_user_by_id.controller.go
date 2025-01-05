package controllers

import (
	"net/http"

	"github.com/HicaroD/learnanything-api/pkg/utils"
	"github.com/labstack/echo/v4"
)

type GetUserByIdRequestBody struct {
	ID   string `validate:"required"`
}

func (h *Handler) GetUserByIdController(ctx echo.Context) error {
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

	// TODO: call service for getting user here

	return ctx.JSON(http.StatusOK, map[string]string{"Hello from": "GetUserByIdController - id: " + req.ID})
}
