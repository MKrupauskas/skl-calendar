package datetime

import (
	"strings"
	"time"
)

func Parse(datetime string) (time.Time, error) {
	loc, err := time.LoadLocation("Europe/Vilnius")
	if err != nil {
			return time.Time{}, err
	}
	trimmed := strings.TrimSpace(datetime)
	date, err := time.ParseInLocation("2006-01-02  15:04", trimmed, loc)
	if err != nil {
			return time.Time{}, err
	}
	return date, nil
}
