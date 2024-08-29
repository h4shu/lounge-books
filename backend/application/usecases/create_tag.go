package usecases

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
)

type CreateTagUsecase interface {
	Execute(context.Context, *inputs.CreateTagInput) (*outputs.CreateTagOutput, error)
}
