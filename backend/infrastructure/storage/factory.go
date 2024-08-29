package storage

import (
	"github.com/h4shu/lounge-books/application/repositories"
)

func NewStorageS3Factory(instance int) repositories.S3 {
	return newS3Handler()
}
