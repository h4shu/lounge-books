package valueobjects

import "errors"

type Day struct {
	i int
}

func NewDay(i int) (*Day, error) {
	if i < 0 || i > 31 {
		return nil, errors.New("invalid day value")
	}
	return &Day{i}, nil
}

func (d *Day) Int() int {
	return d.i
}

func (d *Day) IsNull() bool {
	return d.i == 0
}

func (d *Day) IntPtr() *int {
	if d.IsNull() {
		return nil
	}
	return &d.i
}
