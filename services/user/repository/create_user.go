package repository

import "github.com/HicaroD/learnanything-api/services/user/entity"

func (r *user) CreateUser(user entity.User) (entity.User, error) {
	// TODO: interact with external services
	return user, nil
}
