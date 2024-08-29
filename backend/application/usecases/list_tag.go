package usecases

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
)

type ListTagUsecase interface {
	Execute(context.Context, *inputs.ListTagInput) (*outputs.ListTagOutput, error)
}
