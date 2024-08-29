package presenters

import "github.com/h4shu/lounge-books/application/outputs"

type (
	ListTagPresenter interface {
		Output(*outputs.ListTagOutput) *listTagResponse
	}
	listTagPresenter struct{}
	tagResponse      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	listTagResponse struct {
		Tags []tagResponse `json:"tags"`
	}
)

func NewListTagPresenter() ListTagPresenter {
	return &listTagPresenter{}
}

func (p *listTagPresenter) Output(o *outputs.ListTagOutput) *listTagResponse {
	var res listTagResponse
	res.Tags = []tagResponse{}
	for _, t := range o.Tags {
		tag := tagResponse{
			ID:   t.ID(),
			Name: t.Name(),
		}
		res.Tags = append(res.Tags, tag)
	}
	return &res
}
