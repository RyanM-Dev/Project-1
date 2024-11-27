package services

import (
	"book-app/internal/domain/entities"
	"book-app/internal/domain/repositories"
	"errors"
	"fmt"
)

var ErrBookNotFound = errors.New("book not found")

type BookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) *BookService {
	return &BookService{bookRepo: bookRepo}
}
func (bs *BookService) CreateBook(book *entities.Book) (string, error) {
	if book.Genre == "" {
		return "", errors.New("book genre is required")
	}
	if book.Author == "" {
		return "", errors.New("book author is required")
	}
	if book.Published == nil {
		return "", errors.New("book author is required")
	}
	if book.Title == "" {
		return "", errors.New("book Title is required")
	}

	bookID, err := bs.bookRepo.CreateBook(book)
	if err != nil {
		return "", fmt.Errorf("falied to create book: %w", err)
	}
	return bookID, err
}

func (bs *BookService) GetBook(bookID string) (*entities.Book, error) {
	book, err := bs.bookRepo.GetBookByID(bookID)
	if errors.Is(err, ErrBookNotFound) {
		return nil, fmt.Errorf("book not found: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("falied to get book: %w", err)
	}
	return book, nil
}

func (bs *BookService) UpdateBook(bookID string, book *entities.Book) error {
	if book.Genre == "" {
		return errors.New("book genre is required")
	}
	if book.Author == "" {
		return errors.New("book author is required")
	}
	if book.Published == nil {
		return errors.New("book author is required")
	}
	if book.Title == "" {
		return errors.New("book Title is required")
	}

	err := bs.bookRepo.UpdateBook(bookID, book)
	if err != nil {
		return fmt.Errorf("falied to update book: %w", err)
	}
	return nil
}

func (bs *BookService) DeleteBook(bookID string) error {
	_, err := bs.bookRepo.GetBookByID(bookID)
	if errors.Is(err, ErrBookNotFound) {
		return fmt.Errorf("book not found: %w", err)
	} else if err != nil {
		return fmt.Errorf("falied to get book: %w", err)
	}
	err = bs.bookRepo.DeleteBook(bookID)
	if err != nil {
		return fmt.Errorf("falied to delete book: %w", err)
	}
	return nil
}

func (bs *BookService) GetBooks() ([]*entities.Book, error) {
	books, err := bs.bookRepo.GetAllBooks()
	if err != nil {
		return nil, fmt.Errorf("falied to get books: %w", err)
	}
	return books, nil

}
