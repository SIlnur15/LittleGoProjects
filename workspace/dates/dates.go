package dates

import "errors"

type Dates struct {
	day   int
	month int
	year  int
}

// it's setter function to set day
func (d *Dates) SetDay(newday int) error {
	if newday < 0 || newday > 31 {
		return errors.New("you entered invalid day's value")
	}
	d.day = newday
	return nil
}

// it's setter function to set month
func (m *Dates) SetMonth(newmonth int) error {
	if newmonth < 0 || newmonth > 13 {
		return errors.New("you entered uncorrect month's value")
	}
	m.month = newmonth
	return nil
}

// it's setter function to set year
func (y *Dates) SetYear(newyear int) error {
	if newyear < 0 {
		return errors.New("lad, please enter correct value")
	}
	y.year = newyear
	return nil
}

// it's getter function for day
func (d *Dates) Day() int {
	return d.day
}

// it's getter function for month
func (m *Dates) Month() int {
	return m.month
}

// it's getter function for year
func (y *Dates) Year() int {
	return y.year
}
