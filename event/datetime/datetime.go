package datetime

import (
	"strings"
	"time"
)

func Parse(datetime string) (time.Time, error) {
	trimmed := strings.TrimSpace(datetime)
	date, err := time.Parse("2006-01-02  15:04", trimmed)
	if err != nil {
			return time.Time{}, err
	}
	return date, nil
}
