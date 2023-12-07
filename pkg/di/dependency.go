package di

import (
	"simple-blog-api-golang/internal/auth"
	"simple-blog-api-golang/internal/post"
)

type Dependency struct {
	PostHandler *post.Handler
	AuthHandler *auth.Handler
}

func NewDependency(postHandler *post.Handler, authHandler *auth.Handler) *Dependency {
	return &Dependency{
		PostHandler: postHandler,
		AuthHandler: authHandler,
	}
}
