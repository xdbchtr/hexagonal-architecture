package ports

import "library-app/internal/core/domain"

type BookService interface {
	CreateBook(book *domain.Book) error
	GetBook(id string) (*domain.Book, error)
	GetAllBooks() ([]*domain.Book, error)
	UpdateBook(id string, book *domain.Book) error
	DeleteBook(id string) error
}
