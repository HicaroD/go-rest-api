package server

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"lego-api-go/internal/controllers"
	"lego-api-go/internal/services"
	"lego-api-go/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB_CONFIG utils.DBConfig = utils.DBConfig{
	Driver: "sqlite3",
	DBName: "local",
}

func registerAllHandlers(e *echo.Echo) error {
	var err error

	registerSwaggerHandler(e)

	dbConfig, err := utils.BuildConnectionString(DB_CONFIG)
	if err != nil {
		return err
	}

	db, err := sqlx.Connect(DB_CONFIG.Driver, dbConfig)
	if err != nil {
		return err
	}
	e.Logger.Printf("successfuly connected to local database: %p\n", db)

	// mongoUri := envloader.GetString("MONGODB_URI", "")
	// mongoDb, err := nrdm.ConnectToMongoDB(mongoUri)
	// if err != nil {
	// 	return err
	// }
	// e.Logger.Printf("successfuly connected to mongodb: %p\n", mongoDb)

	loginHandler := &controllers.LoginHandler{}
	loginHandler.RegisterControllers("/login", e)

	userService := services.NewUserService(db)
	userHandler := &controllers.UserHandler{UserService: userService}
	userHandler.RegisterControllers("/users", e)

	return nil
}

func registerSwaggerHandler(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
