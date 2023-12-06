package post

import "context"

type Service interface {
	GetPosts(ctx context.Context) ([]Post, error)
	GetPostById(ctx context.Context, id int) (*Post, error)
}
