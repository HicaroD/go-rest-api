package business

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// TODO: I need a more convenient way of
// converting this struct into a map
func (u User) ToJson() map[string]any {
	return map[string]any{
		"id":        u.ID,
		"name":      u.Name,
		"last_name": u.LastName,
	}
}
