package presenters

import "github.com/h4shu/lounge-books/application/outputs"

type (
	ListBookPresenter interface {
		Output(*outputs.ListBookOutput) *listBookResponse
	}
	listBookPresenterImpl struct{}
	bookResponse          struct {
		ID          int    `json:"id"`
		ISBN        string `json:"isbn"`
		Title       string `json:"title"`
		Description string `json:"description"`
		CoverLink   string `json:"cover_link"`
		PublishedAt string `json:"published_at"`
		Author      string `json:"author"`
		Publisher   string `json:"publisher"`
		PageCount   int    `json:"page_count"`
		DeletedAt   string `json:"deleted_at"`
	}
	listBookResponse struct {
		Books []bookResponse `json:"books"`
	}
)

func NewListBookPresenter() ListBookPresenter {
	return &listBookPresenterImpl{}
}

func (p *listBookPresenterImpl) Output(o *outputs.ListBookOutput) *listBookResponse {
	var res listBookResponse
	res.Books = []bookResponse{}
	for _, b := range o.Books {
		book := bookResponse{
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
		res.Books = append(res.Books, book)
	}
	return &res
}
