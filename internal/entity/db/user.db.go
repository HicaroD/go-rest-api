package db

import (
	"time"

	"gorm.io/gorm"

	"lego-api-go/internal/entity/business"
)

type User struct {
	gorm.Model
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `db:"name"`
	LastName  string `db:"last_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) ToBusiness() business.User {
	return business.User{
		ID:       u.ID,
		Name:     u.Name,
		LastName: u.LastName,
	}
}
