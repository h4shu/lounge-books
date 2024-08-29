package inputs

import (
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type CreateBookInput struct {
	ISBN        *valueobjects.ISBN
	Title       string
	Description string
	CoverLink   string
	PublishedAt *valueobjects.PublishedAt
	Author      *valueobjects.Author
	Publisher   string
	PageCount   int
	TagIDs      []int
}
