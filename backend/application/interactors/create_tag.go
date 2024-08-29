package interactors

import (
	"context"
	"time"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/application/usecases"
)

type CreateTagInteractor struct {
	repository repositories.TagRepository
	ctxTimeout time.Duration
}

func NewCreateTagInteractor(r repositories.TagRepository, t time.Duration) usecases.CreateTagUsecase {
	return &CreateTagInteractor{
		repository: r,
		ctxTimeout: t,
	}
}

func (i *CreateTagInteractor) Execute(ctx context.Context, input *inputs.CreateTagInput) (*outputs.CreateTagOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	err := i.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return &outputs.CreateTagOutput{}, nil
}
