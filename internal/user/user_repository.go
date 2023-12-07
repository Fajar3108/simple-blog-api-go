package user

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindById(ctx context.Context, id int) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Store(ctx context.Context, user *User) error
}
