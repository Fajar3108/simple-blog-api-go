//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"simple-blog-api-golang/internal/post"
)

func InitializeApp() *post.PostHandler {
	wire.Build(post.NewPostHandler, post.NewPostService, post.NewPostRepository)

	return &post.PostHandler{}
}
