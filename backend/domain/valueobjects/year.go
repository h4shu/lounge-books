package valueobjects

import "errors"

type Year struct {
	i int
}

func NewYear(i int) (*Year, error) {
	if i < 0 || i > 9999 {
		return nil, errors.New("invalid year value")
	}
	return &Year{i}, nil
}

func (y *Year) Int() int {
	return y.i
}

func (y *Year) IsNull() bool {
	return y.i == 0
}

func (y *Year) IntPtr() *int {
	if y.IsNull() {
		return nil
	}
	return &y.i
}
