package auth

import "simple-blog-api-golang/internal/user"

type authService struct {
	userRepository user.Repository
}

func NewAuthService(userRepository user.Repository) AuthService {
	return &authService{userRepository: userRepository}
}

func (as authService) Signin() {
	//TODO implement me
	panic("implement me")
}

func (as authService) Signup() {
	//TODO implement me
	panic("implement me")
}

func (as authService) Signout() {
	//TODO implement me
	panic("implement me")
}
