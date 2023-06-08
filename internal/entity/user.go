package entity

type Author struct {
	Base
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Link  *string `json:"link"`
	Admin bool `json:"admin"`
}
