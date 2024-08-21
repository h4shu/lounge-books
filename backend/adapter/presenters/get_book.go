package presenters

import "github.com/h4shu/lounge-books/application/outputs"

type (
	GetBookPresenter interface {
		Output(*outputs.GetBookOutput) *getBookResponse
	}
	getBookPresenterImpl struct{}
	getBookResponse      struct {
		ID          int    `json:"id,string"`
		ISBN        string `json:"isbn"`
		Title       string `json:"title"`
		Description string `json:"description"`
		CoverLink   string `json:"cover_link"`
		PublishedAt string `json:"published_at"`
		Author      string `json:"author"`
		Publisher   string `json:"publisher"`
		PageCount   int    `json:"page_count,string"`
		DeletedAt   string `json:"deleted_at"`
	}
)

func NewGetBookPresenter() GetBookPresenter {
	return &getBookPresenterImpl{}
}

func (p *getBookPresenterImpl) Output(o *outputs.GetBookOutput) *getBookResponse {
	var res getBookResponse
	b := o.Book
	if b != nil {
		res = getBookResponse{
			ID:          b.ID().Int(),
			ISBN:        b.ISBN().String(),
			Title:       b.Title(),
			Description: b.Description(),
			CoverLink:   b.CoverLink(),
			PublishedAt: b.PublishedAt().String(),
			Author:      b.Author().String(),
			Publisher:   b.Publisher(),
			PageCount:   b.PageCount(),
			DeletedAt:   b.DeletedAt().String(),
		}
	}
	return &res
}
