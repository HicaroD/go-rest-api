package users

import (
	"lego-api-go/internal/entity/business"
	"lego-api-go/internal/entity/db"
)

func (s *user) CreateUser(user business.User) (business.User, error) {
	var err error

	newUser := &db.User{Name: user.Name, LastName: user.LastName}
	result := s.DB().Create(newUser)
	err = result.Error
	if err != nil {
		return business.User{}, err
	}

	return newUser.ToBusiness(), nil
}
