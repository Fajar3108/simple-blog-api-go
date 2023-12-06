package user

import (
	"context"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindAll(ctx context.Context) ([]User, error) {
	var users []User

	result := ur.db.WithContext(ctx).Order("id desc").Take(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (ur *userRepository) FindById(ctx context.Context, id int) (*User, error) {
	var user *User

	result := ur.db.WithContext(ctx).Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user *User

	result := ur.db.WithContext(ctx).Where("email=" + email).First(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
