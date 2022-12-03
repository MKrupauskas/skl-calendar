package event

import "time"

type Event struct {
	DateTime time.Time
	Title string
	Location string
}
