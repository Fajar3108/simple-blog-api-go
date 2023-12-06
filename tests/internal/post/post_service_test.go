package post

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"simple-blog-api-golang/internal/post"
	"testing"
)

var postRepository = &RepositoryMock{Mock: mock.Mock{}}
var postService = post.NewPostService(postRepository)

// TestMain sets up the environment variables before running tests
func TestMain(m *testing.M) {
	// Load the environment variables from the .env.test file
	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env.test file:" + err.Error())
	}

	// Run the tests
	exitCode := m.Run()

	// Clean up or perform teardown logic if needed

	// Exit with the test result
	os.Exit(exitCode)
}

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
