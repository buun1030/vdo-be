package api

import (
	"errors"
	"time"
	"vdo-be/pkg/format"
)

// Date is NOT contain only date, but also the timestamp.
// Its purpose is to be overwritten by UnmarshalJSON, MarshalJSON and DecodeRPC.
type Date struct {
	time.Time
}

func NewDate(t time.Time) Date {
	return Date{
		t,
	}
}

func (u *Date) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	// Check for empty string
	if string(data) == `""` {
		return nil
	}

	// data typically looks like `"2012-12-31"`, therefore strip the quotes.
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("invalid date format")
	}
	text := data[1 : len(data)-1]

	t, err := format.ParseDateTime(string(text))
	if err != nil {
		return err
	}

	u.Time = t
	return nil
}

func (u *Date) MarshalJSON() ([]byte, error) {
	dateStr := u.Format(format.DATE_TIME_LAYOUT)
	return []byte("\"" + dateStr + "\""), nil
}

func (u *Date) DecodeRPC(data []byte) error {
	t, err := format.ParseDateTime(string(data))
	if err != nil {
		return err
	}

	u.Time = t
	return nil
}

func (d Date) StartOfMonth() Date {
	return NewDate(time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location()))
}
