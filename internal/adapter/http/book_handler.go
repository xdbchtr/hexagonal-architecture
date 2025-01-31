package http

import (
	"library-app/internal/core/domain"
	"library-app/internal/core/ports"
	"library-app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService ports.BookService
}

func NewBookHandler(bookService ports.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	if err := h.bookService.CreateBook(&book); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create book", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, "Book created successfully", book)
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := h.bookService.GetBook(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Book not found", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Book retrieved successfully", book)
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve books", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Books retrieved successfully", books)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}
	if err := h.bookService.UpdateBook(id, &book); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update book", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Book updated successfully", book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := h.bookService.DeleteBook(id); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete book", err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Book deleted successfully", nil)
}
