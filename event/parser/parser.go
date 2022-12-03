package parser

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
	"mkrup.com/skl-calendar/event"
	"mkrup.com/skl-calendar/event/datetime"
)



func Parse(reader io.ReadCloser) ([]event.Event, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
  if err != nil {
    return nil, err
  }
	// TODO: this is not very robust
	eventsSel := doc.Find("#nav-schedule .card-body")
	events := make([]event.Event, eventsSel.Length())
	eventsSel.Each(func(i int, s *goquery.Selection) {
		dateTime := s.Find(".card-text").Text()
		title := s.Find(".card-title").Text()
		location := s.Find(".date").Text()
		date, err := datetime.Parse(dateTime)
		// TODO: come up with a way to bubble this error in a nicer way
		if err != nil {
				fmt.Printf("parse date: %s\n", err.Error())
				return
		}
		events[i] = event.Event{
			DateTime: date,
			Title:    title,
			Location: location,
		}
	})
	return events, nil
}
