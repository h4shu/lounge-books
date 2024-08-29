package repositories

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/domain/entities"
)

type TagRepository interface {
	Create(context.Context, *inputs.CreateTagInput) error
	FindAll(context.Context) ([]entities.Tag, error)
	FindByID(context.Context, int) (*entities.Tag, error)
}
