package users

import (
	"lego-api-go/internal/entity/business"
	"github.com/Viventio/legos/rdm"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserById(id int) (*business.User, bool, error)
	CreateUser(user business.User) (business.User, error)
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
