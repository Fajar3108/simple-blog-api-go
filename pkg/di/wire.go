//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"simple-blog-api-golang/internal/auth"
	"simple-blog-api-golang/internal/post"
	"simple-blog-api-golang/internal/user"
)

func InitializeApp(db *gorm.DB) (*Dependency, error) {
	wire.Build(
		NewDependency,
		auth.NewAuthHandler,
		post.NewPostHandler,
		post.NewPostService,
		post.NewPostRepository,
		auth.NewAuthService,
		user.NewUserRepository)

	return nil, nil
}
