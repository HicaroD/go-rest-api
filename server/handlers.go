package server

import (
	"github.com/Viventio/legos/rdm"
	"lego-api-go/internal/controllers/login"
	usersH "lego-api-go/internal/controllers/users"
	"lego-api-go/internal/entity/db"
	usersS "lego-api-go/internal/services/users"

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

	// mongoUri := envloader.GetString("MONGODB_URI", "")
	// mongoDb, err := nrdm.ConnectToMongoDB(mongoUri)
	// if err != nil {
	// 	return err
	// }
	// e.Logger.Printf("successfuly connected to mongodb: %p\n", mongoDb)

	loginHandler := &login.Handler{}
	loginHandler.RegisterControllers("/login", e)

	userService := usersS.NewService(localDb)
	userHandler := &usersH.Handler{UserService: userService}
	userHandler.RegisterControllers("/users", e)

	return nil
}
