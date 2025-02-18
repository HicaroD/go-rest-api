package entity

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" db:"name"`
	LastName string `json:"last_name" db:"last_name"`
}
