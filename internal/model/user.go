package model

type User struct {
	BaseModel

	Username    string `gorm:"uniqueIndex" json:"username"`
	Nickname    string `json:"nickname"`
	Email       string `gorm:"uniqueIndex" json:"email"`
	HashedEmail string `gorm:"uniqueIndex" json:"hashed_email"`
	Website     string `json:"website"`
	IsAdmin     *bool  `json:"is_admin"`
	BadgeName   string `json:"badge_name"`
	BadgeColor  string `json:"badge_color"`
	Avatar      string `json:"avatar"` // 头像 URL, 为空时从 Gravatar 获取
	Password    string `json:"password"`
}

func (User) TableName() string {
	return "users"
}
