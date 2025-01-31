package ports

import "library-app/internal/core/domain"

type UserService interface {
	CreateUser(user *domain.User) error
	GetUser(id string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	UpdateUser(id string, user *domain.User) error
	DeleteUser(id string) error
}
