package entity

type Author struct {
	Base
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Link  string `json:"link"`
	IsAdmin bool `json:"is_admin"`
}
