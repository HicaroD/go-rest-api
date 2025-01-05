package main

import (
	"fmt"

	"github.com/HicaroD/learnanything-api/config"
	"github.com/labstack/echo/v4"
)

func main() {
	Start()
}

func Start() {
	cfg := config.InitEnvConfig()

	e := echo.New()
	registerDefaultMiddlewares(e)
	registerAllHandlers(e)

	addr := fmt.Sprintf(":%v", cfg.Port)
	e.Logger.Fatal(e.Start(addr))
}
