package models

import (
	"lego-api-go/internal/user/entity"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Name     string `db:"name"`
	LastName string `db:"last_name"`
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:       u.ID,
		Name:     u.Name,
		LastName: u.LastName,
	}
}
