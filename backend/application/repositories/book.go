package repositories

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type BookRepository interface {
	Create(context.Context, *inputs.CreateBookInput) error
	FindAll(context.Context) ([]entities.Book, error)
	FindByID(context.Context, *valueobjects.BookID) (*entities.Book, error)
	FindByKeywordContaining(context.Context, string) ([]entities.Book, error)
	Update(context.Context, *entities.Book) error
}
