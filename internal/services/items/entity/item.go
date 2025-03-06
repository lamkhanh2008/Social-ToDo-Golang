package entity

import core "social_todo/internal/common"

type TodoItem struct {
	core.SqlModel
	UserID      int          `json:"-" gorm:"column:user_id;"`
	Title       string       `json:"title" gorm:"column:title;"`
	Description string       `json:"description" gorm:"column:description;"`
	Status      string       `json:"status" gorm:"column:status;"`
	LikedCount  int          `json:"liked_count" gorm:"column:liked_count;"`
	Image       *core.Images `json:"images" gorm:"column:images`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	ID          int          `json:"id" gorm:"column:id;"`
	UserID      int          `json:"-" gorm:"column:"user_id;"`
	Title       string       `json:"title" gorm:"column:"title;"`
	Description string       `json:"description" gorm:"column:"description;"`
	Image       *core.Images `json:"images" gorm:"column:"images;"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}
