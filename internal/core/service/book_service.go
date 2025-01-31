package service

import (
	"library-app/internal/core/domain"
	"library-app/internal/core/ports"
)

type bookService struct {
	repo ports.BookRepository
}

func NewBookService(repo ports.BookRepository) ports.BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(book *domain.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) GetBook(id string) (*domain.Book, error) {
	return s.repo.GetByID(id)
}

func (s *bookService) GetAllBooks() ([]*domain.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) UpdateBook(id string, book *domain.Book) error {
	return s.repo.Update(id, book)
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}
