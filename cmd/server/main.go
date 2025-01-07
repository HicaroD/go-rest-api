package main

import (
	"fmt"

	"github.com/HicaroD/api/config"
	"github.com/labstack/echo/v4"
)

func main() {
	Start()
}

func Start() {
	cfg := config.InitEnvConfig()

	e := echo.New()

	registerDefaultMiddlewares(e)

	err := registerAllHandlers(e)
	if err != nil {
		e.Logger.Fatal(err)
	}

	addr := fmt.Sprintf(":%v", cfg.Port)
	e.Logger.Fatal(e.Start(addr))
}
