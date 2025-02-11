package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" db:"name"`
	LastName string `json:"last_name" db:"last_name"`
}
