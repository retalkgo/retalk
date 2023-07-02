package entity

import "time"

// 评论模型
type Comment struct {
	Base
	AuthorID uint   `json:"author_id"`
	Body     string `json:"body"`
	Path     string `json:"path"`
}

type CookedComment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Author    Author    `json:"author"`
	Body      string    `json:"body"`
	Path      string    `json:"path"`
}
