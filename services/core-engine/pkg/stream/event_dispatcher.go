package stream

import (
	"sync"
	"time"
)

// Event represents a system event envelope.
type Event struct {
	ID        string    `json:"id"`
	Topic     string    `json:"topic"`
	Payload   string    `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

// Handler handles incoming events.
type Handler func(event Event)

// Dispatcher manages pub/sub event subscriptions.
type Dispatcher struct {
	mu          sync.RWMutex
	subscribers map[string][]Handler
}

// NewDispatcher creates a new thread-safe event dispatcher.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		subscribers: make(map[string][]Handler),
	}
}

// Subscribe registers a handler for a topic.
func (d *Dispatcher) Subscribe(topic string, h Handler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.subscribers[topic] = append(d.subscribers[topic], h)
}

// Publish distributes an event to all registered topic listeners.
func (d *Dispatcher) Publish(event Event) int {
	d.mu.RLock()
	handlers, exists := d.subscribers[event.Topic]
	d.mu.RUnlock()

	if !exists {
		return 0
	}

	for _, h := range handlers {
		go h(event)
	}
	return len(handlers)
}
