package model

type Site struct {
	BaseModel

	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Domain      string `json:"domain"`
}

func (Site) TableName() string {
	return "sites"
}
