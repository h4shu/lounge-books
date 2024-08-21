package entities

import (
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type (
	Book struct {
		id          *valueobjects.BookID
		isbn        *valueobjects.ISBN
		title       string
		description string
		coverLink   string
		publishedAt *valueobjects.PublishedAt
		author      *valueobjects.Author
		publisher   string
		pageCount   int
		deletedAt   *valueobjects.DeletedAt
	}
)

func NewBook(
	id *valueobjects.BookID,
	isbn *valueobjects.ISBN,
	title string,
	description string,
	coverLink string,
	publishedAt *valueobjects.PublishedAt,
	author *valueobjects.Author,
	publisher string,
	pageCount int,
	deletedAt *valueobjects.DeletedAt,
) *Book {
	return &Book{
		id:          id,
		isbn:        isbn,
		title:       title,
		description: description,
		coverLink:   coverLink,
		publishedAt: publishedAt,
		author:      author,
		publisher:   publisher,
		pageCount:   pageCount,
		deletedAt:   deletedAt,
	}
}

func (b *Book) ID() *valueobjects.BookID {
	return b.id
}

func (b *Book) ISBN() *valueobjects.ISBN {
	return b.isbn
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Description() string {
	return b.description
}

func (b *Book) CoverLink() string {
	return b.coverLink
}

func (b *Book) PublishedAt() *valueobjects.PublishedAt {
	return b.publishedAt
}

func (b *Book) Author() *valueobjects.Author {
	return b.author
}

func (b *Book) Publisher() string {
	return b.publisher
}

func (b *Book) PageCount() int {
	return b.pageCount
}

func (b *Book) DeletedAt() *valueobjects.DeletedAt {
	return b.deletedAt
}
