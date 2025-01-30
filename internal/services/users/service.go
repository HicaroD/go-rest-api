package users

import (
	"lego-api-go/internal/entity/business"
	"lego-api-go/pkg/rdm"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserById(id uint) (*business.User, bool, error)
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
