package valueobjects

import "errors"

type Month struct {
	i int
}

func NewMonth(i int) (*Month, error) {
	if i < 0 || i > 12 {
		return nil, errors.New("invalid month value")
	}
	return &Month{i}, nil
}

func (m *Month) Int() int {
	return m.i
}

func (m *Month) IsNull() bool {
	return m.i == 0
}

func (m *Month) IntPtr() *int {
	if m.IsNull() {
		return nil
	}
	return &m.i
}
