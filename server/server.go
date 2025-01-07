package server

import (
	"fmt"

	"github.com/HicaroD/api/config"
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
	return e.Start(addr)
}
