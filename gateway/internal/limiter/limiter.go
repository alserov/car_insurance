package limiter

import (
	"context"
	"time"
)

type Limiter interface {
	Limit(ctx context.Context) bool
}

const (
	LeakyBucket = iota
)

const (
	DefaultLimit = 10_000
)

const (
	refresh = time.Millisecond * 5
)

func NewLimiter(limiterType, limit int) Limiter {
	switch limiterType {
	case LeakyBucket:
		return newLeakyBucket(limit)
	default:
		panic("invalid limiter type")
	}
}
