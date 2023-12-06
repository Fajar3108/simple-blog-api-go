package post

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Post, error)
	FindById(ctx context.Context, id int) (*Post, error)
}
