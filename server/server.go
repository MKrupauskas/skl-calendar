package server

import (
	"net/http"

	"mkrup.com/skl-calendar/cache"
	"mkrup.com/skl-calendar/event/ics"
)

func Serve(address string, cache *cache.Cache) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/events", events(cache))
	return http.ListenAndServe(address, mux)
}

func events(cache *cache.Cache) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "text/calendar")
			w.Header().Set("charset", "utf-8")
			w.Header().Set("Content-Disposition", "inline")
			w.Header().Set("filename", "calendar.ics")

			// TODO: handle error
			ics.New(cache.Events()).SerializeTo(w)
	})
}

