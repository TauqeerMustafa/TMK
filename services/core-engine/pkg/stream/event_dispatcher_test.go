package stream

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestDispatcherPubSub(t *testing.T) {
	d := NewDispatcher()
	var received atomic.Int32

	d.Subscribe("telemetry.metrics", func(ev Event) {
		received.Add(1)
	})

	count := d.Publish(Event{
		ID:        "evt-001",
		Topic:     "telemetry.metrics",
		Payload:   "cpu_load:0.42",
		Timestamp: time.Now(),
	})

	if count != 1 {
		t.Fatalf("expected 1 handler called, got %d", count)
	}

	time.Sleep(50 * time.Millisecond)
	if received.Load() != 1 {
		t.Errorf("expected received=1, got %d", received.Load())
	}
}
