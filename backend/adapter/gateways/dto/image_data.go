package dto

import (
	"image"
)

type imageDataDTO struct {
	imageDTO *imageDTO
	data     image.Image
}

func NewImageDataDTO(imageDTO *imageDTO, data image.Image) *imageDataDTO {
	return &imageDataDTO{
		imageDTO: imageDTO,
		data:     data,
	}
}

func (d *imageDataDTO) ImageDTO() *imageDTO {
	return d.imageDTO
}

func (d *imageDataDTO) Data() image.Image {
	return d.data
}
