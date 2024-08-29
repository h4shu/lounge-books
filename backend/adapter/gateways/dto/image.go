package dto

import (
	"fmt"

	"github.com/h4shu/lounge-books/domain/entities"
)

const noISBNS3Key = "NoISBN"

type imageDTO struct {
	id          string
	isbn        string
	title       string
	s3Key       string
	format      string
	originalURL string
}

func NewImageDTO(image *entities.Image) *imageDTO {
	var isbnS3Key string
	if image.ISBN().IsNull() {
		isbnS3Key = noISBNS3Key
	} else {
		isbnS3Key = image.ISBN().String()
	}
	s3Key := fmt.Sprintf("%d-%s-%s/%s%s", image.BookID().Int(), isbnS3Key, image.Title(), image.ID().String(), image.Format().FilenameExtension())
	return &imageDTO{
		id:          image.ID().String(),
		isbn:        image.ISBN().String(),
		title:       image.Title(),
		s3Key:       s3Key,
		format:      image.Format().String(),
		originalURL: image.OriginalURL(),
	}
}

func (d *imageDTO) ID() string {
	return d.id
}

func (d *imageDTO) ISBN() string {
	return d.isbn
}

func (d *imageDTO) Title() string {
	return d.title
}

func (d *imageDTO) S3Key() string {
	return d.s3Key
}

func (d *imageDTO) Format() string {
	return d.format
}

func (d *imageDTO) OriginalURL() string {
	return d.originalURL
}
