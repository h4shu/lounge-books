package valueobjects

type Author struct {
	s string
}

func NewAuthor(s string) *Author {
	return &Author{s}
}

func (a *Author) String() string {
	return a.s
}
