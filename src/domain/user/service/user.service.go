package service

import (
	repo "init/project/src/domain/user/repositories"
	"os/user"
)

type UserService interface {
	GetAllUsers() ([]user.User, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (s *userService) GetAllUsers() ([]user.User, error) {
	return s.userRepo.GetAllUsers()
}
