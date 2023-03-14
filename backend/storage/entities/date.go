package entities

import (
	"database/sql/driver"
	"time"
)

type Date struct {
	Day   uint8  `json:"day"`
	Month uint8  `json:"month"`
	Year  uint16 `json:"year"`
}

func Today() Date {
	return DateFromTimestamp(time.Now())
}

func DateFromTimestamp(time time.Time) Date {
	return Date{
		Day:   uint8(time.Day()),
		Month: uint8(time.Month()),
		Year:  uint16(time.Year()),
	}
}

func (date *Date) Scan(src any) error {
	unix := src.(int64)
	timestamp := time.Unix(unix, 0)

	*date = DateFromTimestamp(timestamp)

	return nil
}

func (date *Date) Value() (driver.Value, error) {
	day, month, year := int(date.Year), time.Month(date.Month), int(date.Day)
	timestamp := time.Date(day, month, year, 0, 0, 0, 0, time.UTC)
	unix := timestamp.Unix()

	return unix, nil
}
