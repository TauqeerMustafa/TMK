package pool

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrPoolClosed   = errors.New("worker pool is closed")
	ErrTaskTimeout  = errors.New("task execution timed out")
	ErrQueueFull    = errors.New("task submission queue is full")
)

// Task represents an executable unit of work.
type Task func(ctx context.Context) error

// Pool manages a concurrent set of goroutine workers.
type Pool struct {
	concurrency int
	taskQueue   chan Task
	wg          sync.WaitGroup
	ctx         context.Context
	cancel      context.CancelFunc
	closed      atomic.Bool
	completed   atomic.Uint64
	failed      atomic.Uint64
}

// NewPool initializes a bounded worker pool.
func NewPool(concurrency, queueSize int) *Pool {
	if concurrency <= 0 {
		concurrency = 4
	}
	if queueSize <= 0 {
		queueSize = 1024
	}

	ctx, cancel := context.WithCancel(context.Background())
	p := &Pool{
		concurrency: concurrency,
		taskQueue:   make(chan Task, queueSize),
		ctx:         ctx,
		cancel:      cancel,
	}

	p.start()
	return p
}

func (p *Pool) start() {
	for i := 0; i < p.concurrency; i++ {
		p.wg.Add(1)
		go func(workerID int) {
			defer p.wg.Done()
			for {
				select {
				case <-p.ctx.Done():
					return
				case task, ok := <-p.taskQueue:
					if !ok {
						return
					}
					if err := task(p.ctx); err != nil {
						p.failed.Add(1)
					} else {
						p.completed.Add(1)
					}
				}
			}
		}(i)
	}
}

// Submit enqueues a new task for execution.
func (p *Pool) Submit(t Task) error {
	if p.closed.Load() {
		return ErrPoolClosed
	}
	select {
	case p.taskQueue <- t:
		return nil
	default:
		return ErrQueueFull
	}
}

// Stats returns pool execution metrics.
func (p *Pool) Stats() (completed, failed uint64) {
	return p.completed.Load(), p.failed.Load()
}

// Close gracefully stops the worker pool.
func (p *Pool) Close() {
	if p.closed.CompareAndSwap(false, true) {
		close(p.taskQueue)
		p.wg.Wait()
		p.cancel()
	}
}
