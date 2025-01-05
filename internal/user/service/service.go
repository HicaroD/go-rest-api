package service

import "github.com/HicaroD/learnanything-api/internal/user/entity"

type User interface {
	CreateUser(user entity.User) (entity.User, error)
}

// TODO: inject any dependency here for the user service, such as database
// connection, caching
type user struct{}

// CreateUser implements UserRepository.
func NewUserService() User {
	return &user{}
}
