package valueobjects

import "time"

type DeletedAt struct {
	t *time.Time
}

func NewDeletedAt(s string) (*DeletedAt, error) {
	if s == "" {
		return &DeletedAt{nil}, nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return &DeletedAt{nil}, err
	}
	return &DeletedAt{&t}, nil
}

func (p *DeletedAt) Time() *time.Time {
	return p.t
}

func (p *DeletedAt) String() string {
	if p.t == nil {
		return ""
	}
	return time.Time(*p.t).Format(time.RFC3339)
}
