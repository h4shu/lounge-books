package inputs

import "github.com/h4shu/lounge-books/domain/valueobjects"

type UploadImageInput struct {
	BookID      *valueobjects.BookID
	ISBN        *valueobjects.ISBN
	Title       string
	Data        *valueobjects.ImageData
	Format      *valueobjects.FileFormat
	OriginalURL string
}
