package ports

import "library-app/internal/core/domain"

type BookRepository interface {
	Create(book *domain.Book) error
	GetByID(id string) (*domain.Book, error)
	GetAll() ([]*domain.Book, error)
	Update(id string, book *domain.Book) error
	Delete(id string) error
}
