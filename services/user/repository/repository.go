package repository

import "github.com/HicaroD/learnanything-api/services/user/entity"

type User interface {
	CreateUser(user entity.User) (entity.User, error)
}

// TODO: inject any dependency here for the user repository
type user struct{}

// CreateUser implements UserRepository.
func NewUserRepository() User {
	return &user{}
}

