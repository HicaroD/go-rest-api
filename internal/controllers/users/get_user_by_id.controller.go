package users

import (
	"net/http"

	"lego-api-go/pkg/validators"

	"github.com/labstack/echo/v4"
)

type GetUserByIdRequestBody struct {
	ID uint `param:"id" validate:"required"`
}

func (h *Handler) GetUserByIdController(ctx echo.Context) error {
	req := GetUserByIdRequestBody{}
	err := validators.ValidateRequest(ctx, &req)
	if err != nil {
		return err
	}

	// TODO: conversion of 'id' field to uint does not work properly
	// If I change to a string, it works.

	// user, found, err := h.UserService.GetUserById(req.ID)
	// if err != nil {
	// 	return err
	// }
	//
	// if !found {
	// 	return ctx.JSON(http.StatusNotFound, map[string]string{"detail": fmt.Sprintf("User with id %d not found", req.ID)})
	// }
	//
	// return ctx.JSON(http.StatusOK, map[string]any{"user": user.ToJson()})

	return ctx.NoContent(http.StatusOK)
}
