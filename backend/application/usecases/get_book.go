package usecases

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
)

type GetBookUsecase interface {
	Execute(context.Context, *inputs.GetBookInput) (*outputs.GetBookOutput, error)
}
