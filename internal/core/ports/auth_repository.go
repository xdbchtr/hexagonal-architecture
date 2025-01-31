package ports

import "library-app/internal/core/domain"

type AuthRepository interface {
	CreateUser(user *domain.User) error
	FindUserByEmail(email string) (*domain.User, error)
}
