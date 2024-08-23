package presenters

import "github.com/h4shu/lounge-books/application/outputs"

type (
	DeleteBookPresenter interface {
		Output(*outputs.DeleteBookOutput) *deleteBookResponse
	}
	deleteBookPresenterImpl struct{}
	deleteBookResponse      struct{}
)

func NewDeleteBookPresenter() DeleteBookPresenter {
	return &deleteBookPresenterImpl{}
}

func (p *deleteBookPresenterImpl) Output(o *outputs.DeleteBookOutput) *deleteBookResponse {
	return &deleteBookResponse{}
}
