package usecases

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
)

type ListBookUsecase interface {
	Execute(context.Context, *inputs.ListBookInput) (*outputs.ListBookOutput, error)
}
