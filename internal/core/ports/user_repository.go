package ports

import "library-app/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Update(id string, user *domain.User) error
	Delete(id string) error
}
