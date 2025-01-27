package validators

import (
	// "fmt"
	// "reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateRequest(ctx echo.Context, req any) error {
	// field, ok := reflect.TypeOf(req).Elem().FieldByName("ID")
	// fmt.Println(field, ok)

	if err := ctx.Bind(req); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return err
	}

	return nil
}
