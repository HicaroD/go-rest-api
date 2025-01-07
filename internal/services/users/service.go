package users

import (
	"github.com/HicaroD/api/internal/entity/business"
	"github.com/HicaroD/api/pkg/rdm"
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
