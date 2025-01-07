package users

import "github.com/HicaroD/api/internal/entity"

func (s *user) CreateUser(user entity.User) (entity.User, error) {
	// TODO: interact with external services, such as databases, external APIs
	// and more
	return user, nil
}
