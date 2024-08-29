package presenters

import "github.com/h4shu/lounge-books/application/outputs"

type (
	CreateTagPresenter interface {
		Output(*outputs.CreateTagOutput) *createTagResponse
	}
	createTagPresenter struct{}
	createTagResponse  struct{}
)

func NewCreateTagPresenter() CreateTagPresenter {
	return &createTagPresenter{}
}

func (p *createTagPresenter) Output(o *outputs.CreateTagOutput) *createTagResponse {
	return &createTagResponse{}
}
