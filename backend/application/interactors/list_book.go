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

type ListBookInteractor struct {
	repository repositories.BookRepository
	ctxTimeout time.Duration
}

func NewListBookInteractor(r repositories.BookRepository, t time.Duration) usecases.ListBookUsecase {
	return &ListBookInteractor{
		repository: r,
		ctxTimeout: t,
	}
}

func (i *ListBookInteractor) Execute(ctx context.Context, input *inputs.ListBookInput) (*outputs.ListBookOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	var books []entities.Book
	var err error
	if len(input.SearchKeyword) == 0 {
		books, err = i.repository.FindAll(ctx)
	} else {
		books, err = i.repository.FindByKeywordContaining(ctx, input.SearchKeyword)
	}
	if err != nil {
		return nil, err
	}
	return &outputs.ListBookOutput{
		Books: books,
	}, nil
}
