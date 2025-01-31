package service

import (
	"library-app/internal/core/domain"
	"library-app/internal/core/ports"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUser(id string) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetAllUsers() ([]*domain.User, error) {
	return s.repo.GetAll()
}

func (s *userService) UpdateUser(id string, user *domain.User) error {
	return s.repo.Update(id, user)
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
