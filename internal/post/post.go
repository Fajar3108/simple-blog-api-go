package post

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int          `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title     string       `json:"title" gorm:"not null"`
	Body      string       `json:"body" gorm:"not null"`
	CreatedAt time.Time    `json:"created_at" gorm:"not null"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
