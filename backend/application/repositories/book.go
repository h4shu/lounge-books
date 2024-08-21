package repositories

import (
	"context"

	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type BookRepository interface {
	Create(context.Context, *entities.Book) error
	FindAll(context.Context) ([]entities.Book, error)
	FindByID(context.Context, *valueobjects.BookID) (*entities.Book, error)
}
