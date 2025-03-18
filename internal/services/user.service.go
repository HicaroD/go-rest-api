package services

import (
	"github.com/Viventio/legos/rdm"
	"lego-api-go/internal/entities"
)

// =======================
//      USER SERVICE
// =======================

type UserService interface {
	GetUserById(id int) (*entities.User, bool, error)
	CreateUser(user entities.User) (*entities.User, error)
}

type user struct {
	localDb *rdm.Database
}

func NewUserService(localDb *rdm.Database) UserService {
	return &user{localDb}
}

func (u *user) DB() {
}

// ===============================
//      GET USER BY ID SERVICE
// ===============================

func (s *user) GetUserById(id int) (*entities.User, bool, error) {
	user := &entities.User{ID: id}
	// found := false
	//
	// result := s.DB().First(&user, id)
	// err := result.Error
	// if err != nil {
	// 	return nil, found, err
	// }
	// if result.RowsAffected == 0 {
	// 	return nil, found, nil
	// }
	//
	// found = true
	// return user, found, nil

	// TODO
	return user, false, nil
}

// ===============================
//      CREATE USER SERVICE
// ===============================

func (s *user) CreateUser(user entities.User) (*entities.User, error) {
	// TODO

	// var err error
	newUser := &entities.User{Name: user.Name, LastName: user.LastName}
	// result := s.DB().Create(newUser)
	// err = result.Error
	// if err != nil {
	// 	return nil, err
	// }

	return newUser, nil
}
