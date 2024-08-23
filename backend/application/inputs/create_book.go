package inputs

import "github.com/h4shu/lounge-books/domain/valueobjects"

type CreateBookInput struct {
	ID          *valueobjects.BookID
	ISBN        *valueobjects.ISBN
	Title       string
	Description string
	CoverLink   string
	PublishedAt *valueobjects.PublishedAt
	Author      *valueobjects.Author
	Publisher   string
	PageCount   int
}
