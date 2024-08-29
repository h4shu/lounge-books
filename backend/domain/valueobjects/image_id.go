package valueobjects

import "github.com/google/uuid"

type ImageID struct {
	s string
}

func NewImageID(s string) *ImageID {
	return &ImageID{s}
}

func GenerateImageID() (*ImageID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return NewImageID(id.String()), nil
}

func (i *ImageID) String() string {
	return i.s
}
