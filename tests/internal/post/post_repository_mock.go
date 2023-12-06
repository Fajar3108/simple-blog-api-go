package post

import (
	"context"
	"github.com/stretchr/testify/mock"
	"simple-blog-api-golang/internal/post"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (pr *RepositoryMock) FindAll(ctx context.Context) ([]post.Post, error) {
	return nil, nil
}

func (pr *RepositoryMock) FindById(ctx context.Context, id int) (*post.Post, error) {
	arguments := pr.Mock.Called(ctx, id)

	if arguments.Get(0) == nil {
		return nil, nil
	}

	data := arguments.Get(0).(*post.Post)

	return data, nil
}
