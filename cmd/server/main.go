package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	Start()
}

func Start() {
	e := echo.New()
	registerDefaultMiddlewares(e)
	registerAllHandlers(e)
	e.Logger.Fatal(e.Start(":8080"))
}
