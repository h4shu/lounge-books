package interactors

import (
	"context"
	"time"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/application/usecases"
)

type GetTagInteractor struct {
	repository repositories.TagRepository
	ctxTimeout time.Duration
}

func NewGetTagInteractor(r repositories.TagRepository, t time.Duration) usecases.GetTagUsecase {
	return &GetTagInteractor{
		repository: r,
		ctxTimeout: t,
	}
}

func (i *GetTagInteractor) Execute(ctx context.Context, input *inputs.GetTagInput) (*outputs.GetTagOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	tag, err := i.repository.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return &outputs.GetTagOutput{
		Tag: tag,
	}, nil
}
