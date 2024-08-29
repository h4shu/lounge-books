package repositories

import (
	"context"

	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type ImageRepository interface {
	Create(context.Context, *entities.Image) error
	FindByBookID(context.Context, *valueobjects.BookID) ([]entities.Image, error)
}
