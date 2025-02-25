// Code generated by error generator; DO NOT EDIT.
// Generated at 2025-02-25T15:15:27.651562
package server

import (
	"github.com/labstack/echo/v4"
)

type HTTPError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
}

func NewHTTPError(code string, message string, status int) *HTTPError {
	return &HTTPError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
	}
}

func (e HTTPError) Error() string {
	return e.Code
}

// WARNING: Don't forget to insert this middleware in the Echo framework
func CustomHTTPErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err == nil {
				return nil
			}

			if httpError, ok := err.(*HTTPError); ok {
				return c.JSON(httpError.HTTPStatus, httpError)
			}

			// Handle other error types
			return err
		}
	}
}

// UserNotFoundErr creates a new HTTPError for user_not_found
func UserNotFoundErr() *HTTPError {
	return NewHTTPError("user_not_found", "User does not exists", 404)
}
