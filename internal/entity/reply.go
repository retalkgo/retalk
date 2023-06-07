package entity

// 回复
type Reply struct {
	Base
	AuthorID  uint   `json:"author_id"`
	Body      string `json:"body"`
	CommentID uint   `json:"comment_id"`
}
