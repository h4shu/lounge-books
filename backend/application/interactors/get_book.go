package interactors

import (
	"context"
	"time"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/application/usecases"
)

type GetBookInteractor struct {
	repository repositories.BookRepository
	ctxTimeout time.Duration
}

func NewGetBookInteractor(r repositories.BookRepository, t time.Duration) usecases.GetBookUsecase {
	return &GetBookInteractor{
		repository: r,
		ctxTimeout: t,
	}
}

func (i *GetBookInteractor) Execute(ctx context.Context, input *inputs.GetBookInput) (*outputs.GetBookOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	book, err := i.repository.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return &outputs.GetBookOutput{
		Book: book,
	}, nil
}
