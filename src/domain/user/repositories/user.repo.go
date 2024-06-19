package repositories

import (
	"init/project/src/infrastructure/db"
	"os/user"
)

type UserRepository interface {
	GetAllUsers() ([]user.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetAllUsers() ([]user.User, error) {
	var users []user.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
