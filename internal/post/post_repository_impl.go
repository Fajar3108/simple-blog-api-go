package post

import (
	"context"
	"gorm.io/gorm"
	"simple-blog-api-golang/internal/database"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository() PostRepository {
	return &postRepository{
		db: database.InitDB(),
	}
}

func (pr *postRepository) FindAll(ctx context.Context) ([]Post, error) {
	var posts []Post

	result := pr.db.Order("id desc").Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}
