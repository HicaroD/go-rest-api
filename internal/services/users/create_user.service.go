package users

import (
	"github.com/HicaroD/api/internal/entity/business"
	"github.com/HicaroD/api/internal/entity/db"
)

func (s *user) CreateUser(user business.User) (business.User, error) {
	var err error

	newUser := &db.User{Name: user.Name, LastName: user.LastName}
	result := s.localDb.Conn.Create(newUser)
	err = result.Error
	if err != nil {
		return business.User{}, err
	}

	return newUser.ToBusiness(), nil
}
