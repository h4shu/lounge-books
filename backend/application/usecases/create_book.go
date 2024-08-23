package usecases

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
)

type CreateBookUsecase interface {
	Execute(context.Context, *inputs.CreateBookInput) (*outputs.CreateBookOutput, error)
}
