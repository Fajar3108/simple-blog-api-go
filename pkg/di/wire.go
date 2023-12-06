//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"simple-blog-api-golang/internal/post"
)

func InitializeApp() *post.Handler {
	wire.Build(post.NewPostHandler, post.NewPostService, post.NewPostRepository)

	return &post.Handler{}
}
