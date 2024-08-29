package interactors

import (
	"context"
	"time"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/application/usecases"
)

type ListTagInteractor struct {
	repository repositories.TagRepository
	ctxTimeout time.Duration
}

func NewListTagInteractor(r repositories.TagRepository, t time.Duration) usecases.ListTagUsecase {
	return &ListTagInteractor{
		repository: r,
		ctxTimeout: t,
	}
}

func (i *ListTagInteractor) Execute(ctx context.Context, input *inputs.ListTagInput) (*outputs.ListTagOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	tags, err := i.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return &outputs.ListTagOutput{
		Tags: tags,
	}, nil
}
