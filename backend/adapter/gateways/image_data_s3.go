package gateways

import (
	"context"
	"io"
	"strings"

	"github.com/h4shu/lounge-books/adapter/gateways/dto"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type ImageDataS3 struct {
	s3               repositories.S3
	putObjectInput   repositories.PutObjectInput
	listObjectsInput repositories.ListObjectsInput
	getObjectInput   repositories.GetObjectInput
}

func NewImageDataS3(
	s3 repositories.S3,
	putObjectInput repositories.PutObjectInput,
	listObjectsInput repositories.ListObjectsInput,
	getObjectInput repositories.GetObjectInput,
) repositories.ImageDataRepository {
	return &ImageDataS3{
		s3:               s3,
		putObjectInput:   putObjectInput,
		listObjectsInput: listObjectsInput,
		getObjectInput:   getObjectInput,
	}
}

func (i *ImageDataS3) Save(ctx context.Context, image *entities.Image) error {
	imageDTO := dto.NewImageDTO(image)
	i.putObjectInput.SetKey(imageDTO.S3Key())

	r, w := io.Pipe()
	defer r.Close()
	defer w.Close()

	err := image.Data().EncodeJpeg(w)
	if err != nil {
		return err
	}
	i.putObjectInput.SetBody(r)
	_, err = i.s3.PutObjectWithContext(ctx, i.putObjectInput)
	return err
}

func (i *ImageDataS3) FindByID(ctx context.Context, imageID *valueobjects.ImageID) (*valueobjects.ImageData, error) {
	lo, err := i.s3.ListObjectsWithContext(ctx, i.listObjectsInput)
	if err != nil {
		return nil, err
	}

	var key string
	for _, v := range lo.Contents() {
		if strings.Contains(v.GetKey(), imageID.String()) {
			key = v.GetKey()
			break
		}
	}
	if len(key) == 0 {
		return nil, entities.ErrImageNotFound
	}

	return i.FindByS3Key(ctx, key)
}

func (i *ImageDataS3) FindByS3Key(ctx context.Context, key string) (*valueobjects.ImageData, error) {
	i.getObjectInput.SetKey(key)
	o, err := i.s3.GetObjectWithContext(ctx, i.getObjectInput)
	if err != nil {
		return nil, err
	}
	data := valueobjects.NewImageData()
	err = data.Decode(o.Body())
	if err != nil {
		return nil, err
	}
	return data, nil
}
