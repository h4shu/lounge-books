package entities

import (
	"errors"

	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type Image struct {
	id          *valueobjects.ImageID
	bookID      *valueobjects.BookID
	isbn        *valueobjects.ISBN
	title       string
	data        *valueobjects.ImageData
	format      *valueobjects.FileFormat
	originalURL string
}

var (
	ErrImageNotFound = errors.New("image not found")
)

func NewImage(
	id *valueobjects.ImageID,
	bookID *valueobjects.BookID,
	isbn *valueobjects.ISBN,
	title string,
	data *valueobjects.ImageData,
	format *valueobjects.FileFormat,
	originalURL string,
) *Image {
	return &Image{
		id:          id,
		bookID:      bookID,
		isbn:        isbn,
		title:       title,
		data:        data,
		format:      format,
		originalURL: originalURL,
	}
}

func (i *Image) ID() *valueobjects.ImageID {
	return i.id
}

func (i *Image) BookID() *valueobjects.BookID {
	return i.bookID
}

func (i *Image) ISBN() *valueobjects.ISBN {
	return i.isbn
}

func (i *Image) Title() string {
	return i.title
}

func (i *Image) Data() *valueobjects.ImageData {
	return i.data
}

func (i *Image) Format() *valueobjects.FileFormat {
	return i.format
}

func (i *Image) OriginalURL() string {
	return i.originalURL
}
