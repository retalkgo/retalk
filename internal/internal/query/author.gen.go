// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"time"
)

const TableNameAuthor = "author"

// Author mapped from table <author>
type Author struct {
	ID        int32     `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Link      string    `gorm:"column:link;not null" json:"link"`
	IsAdmin   int32     `gorm:"column:is_admin;not null" json:"is_admin"`
}

// TableName Author's table name
func (*Author) TableName() string {
	return TableNameAuthor
}