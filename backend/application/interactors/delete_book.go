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

type DeleteBookInteractor struct {
	repository repositories.BookRepository
	ctxTimeout time.Duration
}

func NewDeleteBookInteractor(r repositories.BookRepository, t time.Duration) usecases.DeleteBookUsecase {
	return &DeleteBookInteractor{
		repository: r,
		ctxTimeout: t,
	}
}

func (i *DeleteBookInteractor) Execute(ctx context.Context, input *inputs.DeleteBookInput) (*outputs.DeleteBookOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	book, err := i.repository.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if !book.DeletedAt().IsNull() {
		return nil, entities.ErrBookAlreadyDeleted
	}
	t := time.Now()
	updated := entities.NewBook(
		book.ID(),
		book.ISBN(),
		book.Title(),
		book.Description(),
		book.CoverLink(),
		book.PublishedAt(),
		book.Author(),
		book.Publisher(),
		book.PageCount(),
		valueobjects.NewDeletedAt(&t),
	)
	err = i.repository.Update(ctx, updated)
	if err != nil {
		return nil, err
	}
	return &outputs.DeleteBookOutput{}, nil
}
