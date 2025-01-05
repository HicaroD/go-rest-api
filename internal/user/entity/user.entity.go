package entity

import "github.com/google/uuid"

type User struct {
	Uuid     uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
}
