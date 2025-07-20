package model

type Comment struct {
	BaseModel

	// 发送者信息 (登录用户)
	UserID *uint `json:"user_id"`
	User   User  `gorm:"foreignKey:UserID" json:"user"`

	// 发送者信息 (游客)
	GuestName    string `json:"guest_name"`
	GuestEmail   string `json:"guest_email"`
	GuestWebsite string `json:"guest_website"`

	// 归属信息
	SiteID uint   `json:"site_id"`
	Site   Site   `gorm:"foreignKey:SiteID" json:"site"`
	Path   string `json:"path"`

	// 评论内容
	Content string `json:"content"`

	// 回复
	ParentCommentID *uint     `gorm:"index" json:"parent_comment_id"`
	Replies         []Comment `gorm:"foreignKey:ParentCommentID" json:"replies"`
}

func (Comment) TableName() string {
	return "comments"
}

type CookedComment struct {
	BaseModel

	UserID *uint `json:"user_id"`
}
