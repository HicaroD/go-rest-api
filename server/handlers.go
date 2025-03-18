package server

import (
	"lego-api-go/internal/controllers"
	"lego-api-go/internal/services"
	"lego-api-go/utils"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var DB_CONFIG utils.DBConfig = utils.DBConfig{
	Driver: "sqlite",
	DBName: "local.db",
}

func registerAllHandlers(e *echo.Echo) error {
	var err error

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
