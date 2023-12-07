package user

import "github.com/form3tech-oss/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	User *User
}

func NewUserClaims(standardClaims jwt.StandardClaims, user *User) *UserClaims {
	return &UserClaims{StandardClaims: standardClaims, User: user}
}
