package ics

import (
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
	"mkrup.com/skl-calendar/event"
)

func New(events []event.Event) *ics.Calendar {
	cal := ics.NewCalendar()
	for _, event := range events {
		e := cal.AddEvent(uuid.New().String())
		e.SetStartAt(event.DateTime)
		e.SetEndAt(event.DateTime.Add(time.Hour*2))
		e.SetSummary(event.Title)
		e.SetLocation(event.Location)
		e.SetDescription(event.Title)
	}
	return cal
}
