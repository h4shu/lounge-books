package valueobjects

import "fmt"

type PublishedAt struct {
	year  *Year
	month *Month
	day   *Day
}

func NewPublishedAt(year int, month int, day int) (*PublishedAt, error) {
	y, err := NewYear(year)
	if err != nil {
		return nil, err
	}
	m, err := NewMonth(month)
	if err != nil {
		return nil, err
	}
	d, err := NewDay(day)
	if err != nil {
		return nil, err
	}
	return &PublishedAt{
		year:  y,
		month: m,
		day:   d,
	}, nil
}

func (p *PublishedAt) Year() *Year {
	return p.year
}

func (p *PublishedAt) Month() *Month {
	return p.month
}

func (p *PublishedAt) Day() *Day {
	return p.day
}

func (p *PublishedAt) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", p.year.Int(), p.month.Int(), p.day.Int())
}
