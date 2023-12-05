package post

import "context"

type PostRepository interface {
	FindAll(ctx context.Context) ([]Post, error)
}
