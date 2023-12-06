package auth

type AuthService interface {
	Signin()
	Signup()
	Signout()
}
