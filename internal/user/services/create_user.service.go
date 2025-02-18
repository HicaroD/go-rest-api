package users

import (
	"lego-api-go/internal/user/entity"
	"lego-api-go/internal/user/models"
)

func (s *user) CreateUser(user entity.User) (*entity.User, error) {
	var err error

	newUser := &models.User{Name: user.Name, LastName: user.LastName}
	result := s.DB().Create(newUser)
	err = result.Error
	if err != nil {
		return nil, err
	}

	return newUser.ToEntity(), nil
}
