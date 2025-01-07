package controllers

import (
	"net/http"

	"github.com/HicaroD/api/pkg/utils"
	"github.com/labstack/echo/v4"
)

type GetUserByIdRequestBody struct {
	ID string `param:"id" validate:"required"`
}

func (h *Handler) GetUserByIdController(ctx echo.Context) error {
	req := &GetUserByIdRequestBody{}
	err := utils.ValidateRequest(ctx, req)
	if err != nil {
		return err
	}

	// TODO: call service for getting user here

	return ctx.JSON(http.StatusOK, map[string]string{"Hello from": "GetUserByIdController - id: " + req.ID})
}
