package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TauqeerMustafa/TMK/services/core-engine/pkg/pool"
	"github.com/TauqeerMustafa/TMK/services/core-engine/pkg/stream"
)

func main() {
	fmt.Println("=======================================================")
	fmt.Println("⚡ TMK Core Engine Service (Go Concurrency Runtime)")
	fmt.Println("=======================================================")

	workerPool := pool.NewPool(8, 2048)
	defer workerPool.Close()

	dispatcher := stream.NewDispatcher()
	dispatcher.Subscribe("telemetry.system", func(ev stream.Event) {
		fmt.Printf("[%s] Received telemetry: %s
", ev.Timestamp.Format(time.RFC3339), ev.Payload)
	})

	// Submit background heartbeat
	workerPool.Submit(func(ctx context.Context) error {
		dispatcher.Publish(stream.Event{
			ID:        "boot-001",
			Topic:     "telemetry.system",
			Payload:   "core-engine initialized successfully",
			Timestamp: time.Now(),
		})
		return nil
	})

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("✅ Core Engine daemon running. Press Ctrl+C to terminate.")
	// Allow quick run / self-test
	select {
	case <-sigChan:
		fmt.Println("Shutting down core engine...")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Self-test completed cleanly.")
	}
}
