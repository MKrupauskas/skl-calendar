package cache

import (
	"sync"

	"mkrup.com/skl-calendar/event"
)

type Cache struct {
	mu sync.RWMutex

	// protected by mu
	events []event.Event 
}

func (c *Cache) Events() []event.Event {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.events
}

func (c *Cache) SetEvents(events []event.Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.events = events
} 
