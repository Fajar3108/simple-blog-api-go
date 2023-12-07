package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint         `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string       `json:"name" gorm:"not null"`
	Email     string       `json:"email" gorm:"not null;unique"`
	Password  string       `json:"-" gorm:"not null"`
	CreatedAt time.Time    `json:"created_at" gorm:"not null"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
