package post

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simple-blog-api-golang/internal/post"
	"testing"
)

var postRepository = &RepositoryMock{Mock: mock.Mock{}}
var postService = post.NewPostService(postRepository)

func TestCategoryService_GetPostById(t *testing.T) {
	postMockReturn := &post.Post{
		ID:    1,
		Title: "Hello World",
		Body:  "Hello world every body",
	}

	postRepository.Mock.On("FindById", context.Background(), 1).Return(postMockReturn)

	item, err := postService.GetPostById(context.Background(), 1)

	assert.NotNil(t, item)
	assert.Nil(t, err)
}
