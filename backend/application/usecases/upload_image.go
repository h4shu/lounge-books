package usecases

import (
	"context"

	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/outputs"
)

type UploadImageUsecase interface {
	Execute(context.Context, *inputs.UploadImageInput) (*outputs.UploadImageOutput, error)
}
