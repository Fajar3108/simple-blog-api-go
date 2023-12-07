package auth

import (
	"context"
	"simple-blog-api-golang/internal/auth/auth_requests"
)

type AuthService interface {
	Signin(ctx context.Context, userRequest *auth_requests.LoginRequest) (string, error)
	Signup(ctx context.Context, registerRequest *auth_requests.RegisterRequest) (string, error)
}
