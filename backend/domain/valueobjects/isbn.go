package valueobjects

type ISBN struct {
	s string
}

func NewISBN(s string) *ISBN {
	return &ISBN{s}
}

func (i *ISBN) String() string {
	return i.s
}
