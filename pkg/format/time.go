package format

import (
	"time"

	"github.com/pkg/errors"
)

const (
	DATE_LAYOUT      = "2006-01-02"
	DATE_TIME_LAYOUT = "2006-01-02T15:04:05.000Z"
	DATE_MAIL_LAYOUT = "02 January 2006"

	// Timezone
	bangkok_tz = "Asia/Bangkok"
)

var BangkokLocation *time.Location

func init() {
	var err error
	BangkokLocation, err = time.LoadLocation(bangkok_tz)
	if err != nil {
		panic(err)
	}
}

func ParseDateTime(s string) (time.Time, error) {
	allowedLayouts := []string{
		DATE_TIME_LAYOUT,
		DATE_LAYOUT,
		time.RFC3339,
	}

	var err error
	var t time.Time
	for _, layout := range allowedLayouts {
		t, err = time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.Wrap(err, "invalid datetime format")
}
