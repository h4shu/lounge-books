package valueobjects

import "time"

type DeletedAt struct {
	t *time.Time
}

func NewDeletedAt(t *time.Time) *DeletedAt {
	return &DeletedAt{t}
}

func NewDeletedAtFromStr(s string) (*DeletedAt, error) {
	if s == "" {
		return &DeletedAt{nil}, nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return &DeletedAt{nil}, err
	}
	return &DeletedAt{&t}, nil
}

func (d *DeletedAt) Time() *time.Time {
	return d.t
}

func (d *DeletedAt) String() string {
	if d.t == nil {
		return ""
	}
	return time.Time(*d.t).Format(time.RFC3339)
}

func (d *DeletedAt) IsNull() bool {
	return d.t == nil
}
