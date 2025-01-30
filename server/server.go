package server

import (
	"fmt"

	"lego-api-go/config"

	"github.com/labstack/echo/v4"
)

func Start() error {
	cfg := config.InitEnvConfig()

	e := echo.New()
	registerDefaultMiddlewares(e)
	err := registerAllHandlers(e)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf(":%v", cfg.Port)
	if err := e.Start(addr); err != nil {
		return err
	}

	return nil
}
