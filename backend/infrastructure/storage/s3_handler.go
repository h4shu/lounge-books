package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/h4shu/lounge-books/application/repositories"
)

type (
	s3Handler struct {
		svc *s3.S3
	}
	getObjectInput struct {
		bucket *string
		key    *string
	}
	getObjectOutput struct {
		body io.ReadCloser
	}
	listObjectsInput struct {
		bucket *string
		prefix *string
	}
	listObjectsOutput struct {
		contents []*object
	}
	object struct {
		key *string
	}
	putObjectInput struct {
		body   io.ReadSeeker
		bucket *string
		key    *string
	}
)

func newS3Handler() repositories.S3 {
	return &s3Handler{
		svc: s3.New(session.Must(session.NewSession())),
	}
}

func (s *s3Handler) GetObjectWithContext(ctx context.Context, input repositories.GetObjectInput) (repositories.GetObjectOutput, error) {
	o, err := s.svc.GetObjectWithContext(ctx, newS3GetObjectInput(input))
	if err != nil {
		return nil, err
	}
	return &getObjectOutput{
		body: o.Body,
	}, nil
}

func (s *s3Handler) ListObjectsWithContext(ctx context.Context, input repositories.ListObjectsInput) (repositories.ListObjectsOutput, error) {
	o, err := s.svc.ListObjectsWithContext(ctx, newS3ListObjectsInput(input))
	if err != nil {
		return nil, err
	}
	s3Contents := o.Contents
	contents := []*object{}
	for _, c := range s3Contents {
		contents = append(contents, newObject(c))
	}
	return &listObjectsOutput{
		contents: contents,
	}, nil
}

func (s *s3Handler) PutObjectWithContext(ctx context.Context, input repositories.PutObjectInput) (repositories.PutObjectOutput, error) {
	return s.svc.PutObjectWithContext(ctx, newS3PutObjectInput(input))
}

func newS3GetObjectInput(input repositories.GetObjectInput) *s3.GetObjectInput {
	return &s3.GetObjectInput{
		Bucket: input.Bucket(),
		Key:    input.Key(),
	}
}

func newGetObjectInput() repositories.GetObjectInput {
	return &getObjectInput{}
}

func (i *getObjectInput) SetBucket(v string) repositories.GetObjectInput {
	i.bucket = aws.String(v)
	return i
}

func (i *getObjectInput) SetKey(v string) repositories.GetObjectInput {
	i.key = aws.String(v)
	return i
}

func (i *getObjectInput) Bucket() *string {
	return i.bucket
}

func (i *getObjectInput) Key() *string {
	return i.key
}

func (o *getObjectOutput) Body() io.ReadCloser {
	return o.body
}

func newS3ListObjectsInput(input repositories.ListObjectsInput) *s3.ListObjectsInput {
	return &s3.ListObjectsInput{
		Bucket: input.Bucket(),
		Prefix: input.Prefix(),
	}
}

func newListObjectsInput() repositories.ListObjectsInput {
	return &listObjectsInput{}
}

func (i *listObjectsInput) SetBucket(v string) repositories.ListObjectsInput {
	i.bucket = aws.String(v)
	return i
}

func (i *listObjectsInput) SetPrefix(v string) repositories.ListObjectsInput {
	i.prefix = aws.String(v)
	return i
}

func (i *listObjectsInput) Bucket() *string {
	return i.bucket
}

func (i *listObjectsInput) Prefix() *string {
	return i.prefix
}

func (o *listObjectsOutput) Contents() []repositories.Object {
	objs := []repositories.Object{}
	for _, v := range o.contents {
		objs = append(objs, v)
	}
	return objs
}

func newObject(o *s3.Object) *object {
	return &object{
		key: o.Key,
	}
}

func (o *object) Key() *string {
	return o.key
}

func (o *object) GetKey() string {
	return aws.StringValue(o.key)
}

func newS3PutObjectInput(input repositories.PutObjectInput) *s3.PutObjectInput {
	return &s3.PutObjectInput{
		Body:   input.Body(),
		Bucket: input.Bucket(),
		Key:    input.Key(),
	}
}

func newPutObjectInput() repositories.PutObjectInput {
	return &putObjectInput{}
}

func (i *putObjectInput) SetBody(v io.Reader) repositories.PutObjectInput {
	i.body = aws.ReadSeekCloser(v)
	return i
}

func (i *putObjectInput) SetBucket(v string) repositories.PutObjectInput {
	i.bucket = aws.String(v)
	return i
}

func (i *putObjectInput) SetKey(v string) repositories.PutObjectInput {
	i.key = aws.String(v)
	return i
}

func (i *putObjectInput) Body() io.ReadSeeker {
	return i.body
}

func (i *putObjectInput) Bucket() *string {
	return i.bucket
}

func (i *putObjectInput) Key() *string {
	return i.key
}
