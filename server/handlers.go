package server

import (
	usersH "lego-api-go/internal/controllers/users"
	"lego-api-go/internal/entity/db"
	usersS "lego-api-go/internal/services/users"
	"lego-api-go/pkg/envloader"
	"lego-api-go/pkg/nrdm"
	"lego-api-go/pkg/rdm"

	"github.com/labstack/echo/v4"
)

// TODO: I need a way to deal with disconnecting stuff before server shutdown
func registerAllHandlers(e *echo.Echo) error {
	var err error

	dbConfig := rdm.InitSqliteConfig("local.db")
	localDb, err := rdm.Connect(dbConfig)
	if err != nil {
		return err
	}
	err = localDb.AutoMigrate(&db.User{})
	if err != nil {
		return err
	}
	e.Logger.Printf("successfuly connected to localDb: %p\n", localDb)

	mongoUri := envloader.GetString("MONGODB_URI", "")
	mongoDb, err := nrdm.ConnectToMongoDB(mongoUri)
	if err != nil {
		return err
	}
	e.Logger.Printf("successfuly connected to mongodb: %p\n", mongoDb)

	userService := usersS.NewService(localDb)
	userHandler := &usersH.Handler{UserService: userService}
	userHandler.RegisterControllers("/users", e)

	return nil
}
