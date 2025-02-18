package users

import (
	"lego-api-go/internal/user/entity"
	"lego-api-go/internal/user/models"
)

func (s *user) GetUserById(id int) (*entity.User, bool, error) {
	user := &models.User{ID: id}
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
	usr := user.ToEntity()
	return usr, found, nil
}
