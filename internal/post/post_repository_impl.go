package post

import (
	"context"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) Repository {
	return &postRepository{db: db}
}

func (pr *postRepository) FindAll(ctx context.Context) ([]Post, error) {
	var posts []Post

	result := pr.db.WithContext(ctx).Order("id desc").Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (pr *postRepository) FindById(ctx context.Context, id int) (*Post, error) {
	var post *Post

	result := pr.db.WithContext(ctx).First(&post, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}
