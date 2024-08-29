package usecases

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
)

type GetTagUsecase interface {
	Execute(context.Context, *inputs.GetTagInput) (*outputs.GetTagOutput, error)
}
