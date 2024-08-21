package interactors

import (
	"context"
	"time"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/application/usecases"
	"github.com/h4shu/lounge-books/domain/entities"
)

type CreateBookInteractor struct {
	repository repositories.BookRepository
	ctxTimeout time.Duration
}

func NewCreateBookInteractor(r repositories.BookRepository, t time.Duration) usecases.CreateBookUsecase {
	return &CreateBookInteractor{
		repository: r,
		ctxTimeout: t,
	}
}

func (i *CreateBookInteractor) Execute(ctx context.Context, input *inputs.CreateBookInput) (*outputs.CreateBookOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	book := entities.NewBook(
		input.ID,
		input.ISBN,
		input.Title,
		input.Description,
		input.CoverLink,
		input.PublishedAt,
		input.Author,
		input.Publisher,
		input.PageCount,
		nil,
	)
	err := i.repository.Create(ctx, book)
	if err != nil {
		return nil, err
	}
	return &outputs.CreateBookOutput{}, nil
}
