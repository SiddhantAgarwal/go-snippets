package tokenbucket

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu           sync.Mutex
	rate         float64
	capacity     float64
	tokens       float64
	lastRefilled time.Time
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefilled).Seconds()
	tb.tokens += tb.rate * elapsed

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	tb.lastRefilled = now

	if tb.tokens >= 1.0 {
		tb.tokens -= 1
		return true
	}

	return false
}

func NewTokenBucket(rate, capacity float64) *TokenBucket {
	return &TokenBucket{
		rate:         rate,
		capacity:     capacity,
		tokens:       capacity,
		lastRefilled: time.Now(),
	}
}
