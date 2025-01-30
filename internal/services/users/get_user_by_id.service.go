package users

import (
	"lego-api-go/internal/entity/business"
	"lego-api-go/internal/entity/db"
)

func (s *user) GetUserById(id int) (*business.User, bool, error) {
	user := &db.User{ID: id}
	found := false

	result := s.DB().First(&user, id)
	err := result.Error
	if err != nil {
		return nil, found, err
	}
	if result.RowsAffected == 0 {
		return nil, found, nil
	}

	found = true
	usr := user.ToBusiness()
	return &usr, found, nil
}
