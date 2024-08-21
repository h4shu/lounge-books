package valueobjects

import "time"

type PublishedAt struct {
	t *time.Time
}

func NewPublishedAt(t *time.Time) *PublishedAt {
	return &PublishedAt{t}
}

func NewPublishedAtFromStr(s string) (*PublishedAt, error) {
	if s == "" {
		return &PublishedAt{nil}, nil
	}
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return &PublishedAt{nil}, err
	}
	return &PublishedAt{&t}, nil
}

func (p *PublishedAt) Time() *time.Time {
	return p.t
}

func (p *PublishedAt) String() string {
	if p.t == nil {
		return ""
	}
	return time.Time(*p.t).Format(time.DateOnly)
}
