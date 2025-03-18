package server

import (
	"lego-api-go/internal/controllers"
	"lego-api-go/internal/services"

	"github.com/Viventio/legos/rdm"
	_ "github.com/Viventio/legos/rdm/drivers/sqlite"

	"github.com/labstack/echo/v4"
)

func registerAllHandlers(e *echo.Echo) error {
	var err error

	dbConfig := rdm.InitSqliteConfig("local.db")
	localDb, err := rdm.Connect(dbConfig)
	if err != nil {
		return err
	}
	// err = localDb.AutoMigrate(&userM.User{})
	// if err != nil {
	// 	return err
	// }
	e.Logger.Printf("successfuly connected to localDb: %p\n", localDb)

	// mongoUri := envloader.GetString("MONGODB_URI", "")
	// mongoDb, err := nrdm.ConnectToMongoDB(mongoUri)
	// if err != nil {
	// 	return err
	// }
	// e.Logger.Printf("successfuly connected to mongodb: %p\n", mongoDb)

	loginHandler := &controllers.LoginHandler{}
	loginHandler.RegisterControllers("/login", e)

	userService := services.NewUserService(localDb)
	userHandler := &controllers.UserHandler{UserService: userService}
	userHandler.RegisterControllers("/users", e)

	return nil
}
