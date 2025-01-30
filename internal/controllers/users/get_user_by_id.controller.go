package users

import (
	"fmt"
	"net/http"
	"strconv"

	"lego-api-go/pkg/validators"

	"github.com/labstack/echo/v4"
)

type GetUserByIdRequestBody struct {
	ID string `param:"id" validate:"required"`
}

func (h *Handler) GetUserByIdController(ctx echo.Context) error {
	req := GetUserByIdRequestBody{}
	err := validators.ValidateRequest(ctx, &req)
	if err != nil {
		return err
	}

	userId, err := strconv.Atoi(req.ID)
	if err != nil {
		return err
	}

	user, found, err := h.UserService.GetUserById(userId)
	if err != nil {
		return err
	}

	if !found {
		return ctx.JSON(http.StatusNotFound, map[string]string{"detail": fmt.Sprintf("User with id %s not found", req.ID)})
	}

	return ctx.JSON(http.StatusOK, map[string]any{"user": user})
}
