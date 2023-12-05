package post

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID    int
	Title string
	Body  string
}
