package limiter

import (
	"context"
	"sync"
	"time"
)

func newLeakyBucket(lim int) Limiter {
	tokens := make(chan struct{}, lim)
	for i := 0; i < lim; i++ {
		tokens <- struct{}{}
	}

	return &leakyBucket{
		tokens: tokens,
	}
}

type leakyBucket struct {
	tokens chan struct{}

	lastTokenTaken time.Time

	mu sync.RWMutex
}

func (l *leakyBucket) Limit(ctx context.Context) bool {
	if time.Since(l.lastTokenTaken) >= refresh {
		select {
		case l.tokens <- struct{}{}:
		default:
		}
	}

	select {
	case <-ctx.Done():
		return true
	case <-l.tokens:
		l.mu.RLock()
		l.lastTokenTaken = time.Now()
		l.mu.RUnlock()

		return false
	}
}
