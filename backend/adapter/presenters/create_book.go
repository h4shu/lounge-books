package presenters

import "github.com/h4shu/lounge-books/application/outputs"

type (
	CreateBookPresenter interface {
		Output(*outputs.CreateBookOutput) *createBookResponse
	}
	createBookPresenterImpl struct{}
	createBookResponse      struct{}
)

func NewCreateBookPresenter() CreateBookPresenter {
	return &createBookPresenterImpl{}
}

func (p *createBookPresenterImpl) Output(o *outputs.CreateBookOutput) *createBookResponse {
	return &createBookResponse{}
}
