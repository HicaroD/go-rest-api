package users

import (
	"net/http"

	"lego-api-go/internal/entity/business"
	"lego-api-go/pkg/validators"

	"github.com/labstack/echo/v4"
)

type CreateUserRequestBody struct {
	Name     string `json:"name"       validate:"required"`
	LastName string `json:"last_name"  validate:"required"`
}

func (r CreateUserRequestBody) toUser() business.User {
	return business.User{
		Name:     r.Name,
		LastName: r.LastName,
	}
}

func (h *Handler) CreateUserController(ctx echo.Context) error {
	req := CreateUserRequestBody{}
	err := validators.ValidateRequest(ctx, &req)
	if err != nil {
		return err
	}

	user := req.toUser()
	newUser, err := h.UserService.CreateUser(user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, map[string]any{"user": newUser})
}
