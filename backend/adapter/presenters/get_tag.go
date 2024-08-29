package presenters

import "github.com/h4shu/lounge-books/application/outputs"

type (
	GetTagPresenter interface {
		Output(*outputs.GetTagOutput) *getTagResponse
	}
	getTagPresenterImpl struct{}
	getTagResponse      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func NewGetTagPresenter() GetTagPresenter {
	return &getTagPresenterImpl{}
}

func (p *getTagPresenterImpl) Output(o *outputs.GetTagOutput) *getTagResponse {
	var res getTagResponse
	t := o.Tag
	if t != nil {
		res = getTagResponse{
			ID:   t.ID(),
			Name: t.Name(),
		}
	}
	return &res
}
