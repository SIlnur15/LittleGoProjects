package calendar

import "errors"

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDay(newday int) error {
	if 0 > newday || newday > 32 {
		return errors.New("this value of day is impossible")
	}
	d.day = newday
	return nil
}

func (m *Date) SetMonth(newmonth int) error {
	if newmonth > 12 || newmonth < 0 {
		return errors.New("oh no this month is impossible")
	}
	m.month = newmonth
	return nil
}

func (y *Date) SetYear(newyear int) error {
	if newyear < 1 {
		return errors.New("that year is impossible")
	}
	y.year = newyear
	return nil
}
