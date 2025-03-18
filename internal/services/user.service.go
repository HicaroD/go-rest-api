package services

import (
	"github.com/jmoiron/sqlx"
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
	db *sqlx.DB
}

func NewUserService(localDb *sqlx.DB) UserService {
	return &user{localDb}
}

// ===============================
//      GET USER BY ID SERVICE
// ===============================

func (s *user) GetUserById(id int) (*entities.User, bool, error) {
	user := new(entities.User)

	query := "SELECT * FROM users WHERE id=?"
	err := s.db.Get(user, query, id)

	return user, err != nil, err
}

// ===============================
//      CREATE USER SERVICE
// ===============================

func (s *user) CreateUser(user entities.User) (*entities.User, error) {
	newUser := &entities.User{Name: user.Name, LastName: user.LastName}

	query := "INSERT INTO users (name, lastName)"
	res := s.db.MustExec(query, user.Name, user.LastName)
	if _, err := res.RowsAffected(); err == nil {
		return nil, err
	}

	return newUser, nil
}
