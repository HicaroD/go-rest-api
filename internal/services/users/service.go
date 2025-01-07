package users

import "github.com/HicaroD/api/internal/entity"

type UserService interface {
	CreateUser(user entity.User) (entity.User, error)
}

// TODO: inject any dependency here for the user service, such as database
// connection, caching
type user struct{}

// CreateUser implements UserRepository.
func NewService() UserService {
	return &user{}
}
