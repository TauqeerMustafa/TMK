package pool

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestWorkerPoolExecution(t *testing.T) {
	p := NewPool(4, 50)
	defer p.Close()

	var counter atomic.Int32
	tasks := 20
	for i := 0; i < tasks; i++ {
		err := p.Submit(func(ctx context.Context) error {
			counter.Add(1)
			return nil
		})
		if err != nil {
			t.Fatalf("unexpected submit error: %v", err)
		}
	}

	time.Sleep(100 * time.Millisecond)
	if counter.Load() != int32(tasks) {
		t.Errorf("expected %d tasks completed, got %d", tasks, counter.Load())
	}

	completed, failed := p.Stats()
	if completed != uint64(tasks) || failed != 0 {
		t.Errorf("unexpected stats: completed=%d, failed=%d", completed, failed)
	}
}
