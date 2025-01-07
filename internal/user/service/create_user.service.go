package service

import "github.com/HicaroD/api/internal/user/entity"

func (r *user) CreateUser(user entity.User) (entity.User, error) {
	// TODO: interact with external services, such as databases, external APIs
	// and more
	return user, nil
}
