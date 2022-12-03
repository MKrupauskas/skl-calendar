package main

import (
	"fmt"
	"os"
	"time"

	"mkrup.com/skl-calendar/cache"
	"mkrup.com/skl-calendar/event/parser"
	"mkrup.com/skl-calendar/fetcher"
	"mkrup.com/skl-calendar/server"
	// "mkrup.com/skl-calendar/event/parser"
	// "mkrup.com/skl-calendar/fetcher"
)

func refresh(c *cache.Cache) error {
	resp, err := fetcher.Fetch("https://www.sostineskl.lt/komandos/stumokliai")
	if err != nil {
			return err
	}
	// TODO: handle error
	defer resp.Body.Close()
	events, err := parser.Parse(resp.Body)
	if err != nil {
			return err
	}
	c.SetEvents(events)
	return nil
}

func periodic(c *cache.Cache) {
	fmt.Println("initial refresh")
	if err := refresh(c); err != nil {
		fmt.Printf("initial refresh: %s\n", err.Error())
	}

	ticker := time.NewTicker(6 * time.Hour)
	defer ticker.Stop()

	for {
			t := <-ticker.C
			fmt.Printf("refresh events: %s\n", t)
			refresh(c)
	}
}

func run() error {
	c := &cache.Cache{}
	go periodic(c)
	
	fmt.Println("serving port http://localhost:8080/events")
	return server.Serve(":8080", c)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)	
	}
}
