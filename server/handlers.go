package server

import (
	usersH "github.com/HicaroD/api/internal/controllers/users"
	"github.com/HicaroD/api/internal/entity/db"
	usersS "github.com/HicaroD/api/internal/services/users"
	"github.com/HicaroD/api/pkg/rdm"
	"github.com/labstack/echo/v4"
)

func registerAllHandlers(e *echo.Echo) error {
	// TODO: initialize databases, loggers, configuration files and more

	dbConfig := rdm.InitSqliteConfig("local.db")
	localDb, err := rdm.Connect(dbConfig)
	if err != nil {
		return err
	}
	err = localDb.AutoMigrate(&db.User{})
	if err != nil {
		return err
	}
	e.Logger.Printf("localDb: %p\n", localDb)

	userService := usersS.NewService(localDb)
	userHandler := &usersH.Handler{UserService: userService}
	userHandler.RegisterControllers("/users", e)

	return nil
}
