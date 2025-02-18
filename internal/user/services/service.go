package users

import (
	"lego-api-go/internal/user/entity"

	"github.com/Viventio/legos/rdm"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserById(id int) (*entity.User, bool, error)
	CreateUser(user entity.User) (*entity.User, error)
}

type user struct {
	localDb *rdm.Database
}

func NewService(localDb *rdm.Database) UserService {
	return &user{localDb}
}

func (u *user) DB() *gorm.DB {
	return u.localDb.Conn
}
