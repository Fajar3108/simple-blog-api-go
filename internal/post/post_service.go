package post

import "context"

type PostService interface {
	GetPosts(ctx context.Context) ([]Post, error)
}
