package auth

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"simple-blog-api-golang/internal/auth/auth_requests"
	"simple-blog-api-golang/internal/user"
	"simple-blog-api-golang/pkg/encryption"
	"simple-blog-api-golang/pkg/jwt_service"
	"simple-blog-api-golang/pkg/validation"
	"time"
)

type authService struct {
	userRepository user.Repository
}

func NewAuthService(userRepository user.Repository) AuthService {
	return &authService{userRepository: userRepository}
}

func (as *authService) Signin(ctx context.Context, userRequest *auth_requests.LoginRequest) (string, error) {
	if err := validation.Validate[*auth_requests.LoginRequest](userRequest); err != nil {
		return "", fmt.Errorf("Errror validation: ", err.(validator.ValidationErrors).Error())
	}
	userByEmail, err := as.userRepository.FindByEmail(ctx, userRequest.Email)

	if err != nil {
		return "", err
	}

	if err := encryption.Check(userRequest.Password, userByEmail.Password); err != nil {
		return "", err
	}

	token, err := jwt_service.GenerateJWT(userByEmail)
	if err != nil {
		return "", err
	}

	return token, err
}

func (as *authService) Signup(ctx context.Context, registerRequest *auth_requests.RegisterRequest) (string, error) {
	if err := validation.Validate[*auth_requests.RegisterRequest](registerRequest); err != nil {
		return "", fmt.Errorf("Errror validation: ", err)
	}

	userByEmail, _ := as.userRepository.FindByEmail(ctx, registerRequest.Email)

	if userByEmail != nil {
		return "", fmt.Errorf("email already exist")
	}

	passwordHashed, err := encryption.Hash(registerRequest.Password)

	if err != nil {
		return "", err
	}

	newUser := user.User{
		Name:      registerRequest.Name,
		Email:     registerRequest.Email,
		Password:  passwordHashed,
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{Valid: false},
		DeletedAt: sql.NullTime{Valid: false},
	}

	if err := as.userRepository.Store(ctx, &newUser); err != nil {
		return "", err
	}

	token, err := jwt_service.GenerateJWT(userByEmail)

	if err != nil {
		return "", nil
	}

	return token, nil
}
