package repositories

import "book-app/internal/domain/entities"

type BookRepository interface {
	CreateBook(book *entities.Book) (string, error)
	GetBookByID(bookID string) (*entities.Book, error)
	GetAllBooks() ([]*entities.Book, error)
	UpdateBook(bookID string, book *entities.Book) error
	DeleteBook(bookID string) error
}
