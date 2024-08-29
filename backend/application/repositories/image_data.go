package repositories

import (
	"context"

	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type ImageDataRepository interface {
	Save(context.Context, *entities.Image) error
	FindByID(context.Context, *valueobjects.ImageID) (*valueobjects.ImageData, error)
	FindByS3Key(context.Context, string) (*valueobjects.ImageData, error)
}
