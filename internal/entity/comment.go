package entity

// 评论模型
type Comment struct {
	Base
	AuthorID uint   `json:"author_id"`
	Body     string `json:"body"`
	Path string `json:"path"`
}
