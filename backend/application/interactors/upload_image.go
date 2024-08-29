package interactors

import (
	"context"
	"time"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/application/usecases"
	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type UploadImageInteractor struct {
	repository repositories.ImageDataRepository
	ctxTimeout time.Duration
}

func NewUploadImageInteractor(r repositories.ImageDataRepository, t time.Duration) usecases.UploadImageUsecase {
	return &UploadImageInteractor{
		repository: r,
		ctxTimeout: t,
	}
}

func (i *UploadImageInteractor) Execute(ctx context.Context, input *inputs.UploadImageInput) (*outputs.UploadImageOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	id, err := valueobjects.GenerateImageID()
	if err != nil {
		return nil, err
	}
	img := entities.NewImage(
		id,
		input.BookID,
		input.ISBN,
		input.Title,
		input.Data,
		input.Format,
		input.OriginalURL,
	)
	err = i.repository.Save(ctx, img)
	if err != nil {
		return nil, err
	}
	return &outputs.UploadImageOutput{
		ID: id,
	}, nil
}
