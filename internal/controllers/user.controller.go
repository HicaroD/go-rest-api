package controllers

import (
	"fmt"
	"lego-api-go/internal/entities"
	"lego-api-go/internal/services"
	"net/http"
	"strconv"

	"github.com/Viventio/legos/jwt"
	"github.com/Viventio/legos/validators"
	"github.com/labstack/echo/v4"
)

// =======================
//      USER HANDLER
// =======================

type UserHandler struct {
	UserService services.UserService
}

func (h *UserHandler) RegisterControllers(prefix string, e *echo.Echo) {
	g := e.Group(prefix)
	g.Use(jwt.AuthMiddleware)

	g.GET("/", h.GetAllUsersController)
	g.GET("/:id", h.GetUserByIdController)
	g.POST("/", h.CreateUserController)
}

// ==============================
//      CREATE USER CONTROLLER
// ==============================

type CreateUserRequestBody struct {
	Name     string `json:"name"       validate:"required"`
	LastName string `json:"last_name"  validate:"required"`
}

func (r CreateUserRequestBody) toUser() entities.User {
	return entities.User{
		Name:     r.Name,
		LastName: r.LastName,
	}
}

func (h *UserHandler) CreateUserController(ctx echo.Context) error {
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

// =================================
//      GET ALL USERS CONTROLLER
// =================================

func (h *UserHandler) GetAllUsersController(ctx echo.Context) error {
	// call service
	return ctx.JSON(http.StatusOK, map[string]string{"Hello from": "GetAllUsersController"})
}

// ==================================
//      GET USER BY ID CONTROLLER
// ==================================

type GetUserByIdRequestBody struct {
	ID string `param:"id" validate:"required"`
}

func (h *UserHandler) GetUserByIdController(ctx echo.Context) error {
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
