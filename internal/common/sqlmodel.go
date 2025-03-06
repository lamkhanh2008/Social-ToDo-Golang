package core

import "time"

type SqlModel struct {
	ID        int        `json:"id" gorm:"column:id;" db:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;" db:"created_at"`
}

func NewSQLModel() SqlModel {
	now := time.Now().UTC()
	return SqlModel{
		ID:        0,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
