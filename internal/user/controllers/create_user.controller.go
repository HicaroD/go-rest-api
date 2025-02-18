package users

import (
	"lego-api-go/internal/user/entity"
	"net/http"

	"github.com/Viventio/legos/validators"
	"github.com/labstack/echo/v4"
)

type CreateUserRequestBody struct {
	Name     string `json:"name"       validate:"required"`
	LastName string `json:"last_name"  validate:"required"`
}

func (r CreateUserRequestBody) toUser() entity.User {
	return entity.User{
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
