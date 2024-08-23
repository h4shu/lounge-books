package valueobjects

type BookID struct {
	i int
}

func NewBookID(i int) *BookID {
	return &BookID{i}
}

func (b *BookID) Int() int {
	return b.i
}
