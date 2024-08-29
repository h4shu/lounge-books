package repositories

import (
	"context"
	"io"
)

type (
	S3 interface {
		GetObjectWithContext(ctx context.Context, input GetObjectInput) (GetObjectOutput, error)
		ListObjectsWithContext(ctx context.Context, input ListObjectsInput) (ListObjectsOutput, error)
		PutObjectWithContext(ctx context.Context, input PutObjectInput) (PutObjectOutput, error)
	}
	GetObjectInput interface {
		SetBucket(v string) GetObjectInput
		SetKey(v string) GetObjectInput
		Bucket() *string
		Key() *string
	}
	GetObjectOutput interface {
		Body() io.ReadCloser
	}
	ListObjectsInput interface {
		SetBucket(v string) ListObjectsInput
		SetPrefix(v string) ListObjectsInput
		Bucket() *string
		Prefix() *string
	}
	ListObjectsOutput interface {
		Contents() []Object
	}
	Object interface {
		Key() *string
		GetKey() string
	}
	PutObjectInput interface {
		SetBody(v io.Reader) PutObjectInput
		SetBucket(v string) PutObjectInput
		SetKey(v string) PutObjectInput
		Body() io.ReadSeeker
		Bucket() *string
		Key() *string
	}
	PutObjectOutput interface{}
)
