package repositories

import (
	"net/url"

	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type Web interface {
	SearchImageURL(isbn *valueobjects.ISBN) (url.URL, error)
	GetImageData(url url.URL)
}
