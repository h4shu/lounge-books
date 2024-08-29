package presenters

import "github.com/h4shu/lounge-books/application/outputs"

type (
	GetBookPresenter interface {
		Output(*outputs.GetBookOutput) *getBookResponse
	}
	getBookPresenterImpl struct{}
	getBookResponse      struct {
		ID             int      `json:"id"`
		ISBN           string   `json:"isbn"`
		Title          string   `json:"title"`
		Description    string   `json:"description"`
		CoverLink      string   `json:"cover_link"`
		PublishedYear  *int     `json:"published_year"`
		PublishedMonth *int     `json:"published_month"`
		PublishedDay   *int     `json:"published_day"`
		Author         string   `json:"author"`
		Publisher      string   `json:"publisher"`
		PageCount      int      `json:"page_count"`
		TagNames       []string `json:"tag_names"`
		DeletedAt      string   `json:"deleted_at"`
	}
)

func NewGetBookPresenter() GetBookPresenter {
	return &getBookPresenterImpl{}
}

func (p *getBookPresenterImpl) Output(o *outputs.GetBookOutput) *getBookResponse {
	var res getBookResponse
	b := o.Book
	if b != nil {
		tagNames := []string{}
		for _, t := range b.Tags() {
			tagNames = append(tagNames, t.Name())
		}
		res = getBookResponse{
			ID:             b.ID().Int(),
			ISBN:           b.ISBN().String(),
			Title:          b.Title(),
			Description:    b.Description(),
			CoverLink:      b.CoverLink(),
			PublishedYear:  b.PublishedAt().Year().IntPtr(),
			PublishedMonth: b.PublishedAt().Month().IntPtr(),
			PublishedDay:   b.PublishedAt().Day().IntPtr(),
			Author:         b.Author().String(),
			Publisher:      b.Publisher(),
			PageCount:      b.PageCount(),
			TagNames:       tagNames,
			DeletedAt:      b.DeletedAt().String(),
		}
	}
	return &res
}
