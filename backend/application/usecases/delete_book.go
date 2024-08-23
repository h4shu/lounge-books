package usecases

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
)

type DeleteBookUsecase interface {
	Execute(context.Context, *inputs.DeleteBookInput) (*outputs.DeleteBookOutput, error)
}
