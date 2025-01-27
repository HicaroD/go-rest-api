package users

import (
	"github.com/HicaroD/api/internal/entity/business"
	"github.com/HicaroD/api/internal/entity/db"
)

func (s *user) GetUserById(id uint) (*business.User, bool, error) {
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
