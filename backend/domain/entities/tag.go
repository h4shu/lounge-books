package entities

import "errors"

type Tag struct {
	id   int
	name string
}

var (
	ErrTagNotFound = errors.New("tag not found")
)

func NewTag(
	id int,
	name string,
) *Tag {
	return &Tag{
		id:   id,
		name: name,
	}
}

func (t *Tag) ID() int {
	return t.id
}

func (t *Tag) Name() string {
	return t.name
}
