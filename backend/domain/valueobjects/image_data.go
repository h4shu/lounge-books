package valueobjects

import (
	"image"
	"image/jpeg"
	"io"
)

type ImageData struct {
	image image.Image
	str   string
}

func NewImageData() *ImageData {
	return &ImageData{}
}

func (i *ImageData) Decode(r io.Reader) (err error) {
	i.image, i.str, err = image.Decode(r)
	return
}

func (i *ImageData) EncodeJpeg(w io.Writer) error {
	return jpeg.Encode(w, i.image, nil)
}
