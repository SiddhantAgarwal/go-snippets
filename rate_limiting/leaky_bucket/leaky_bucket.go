package leakybucket

import (
	"math"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity     float64
	leakRate     float64
	usedCapacity float64
	lastLeakTime time.Time
	mu           sync.Mutex
}

func NewLeakyBucket(capacity, leakRate float64) *LeakyBucket {
	return &LeakyBucket{
		capacity:     capacity,
		leakRate:     leakRate,
		lastLeakTime: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(lb.lastLeakTime).Seconds()
	leakedAmount := elapsed * lb.leakRate

	lb.usedCapacity = math.Max(0, lb.usedCapacity-leakedAmount)

	lb.lastLeakTime = now
	if lb.usedCapacity < lb.capacity {
		lb.usedCapacity++
		return true
	}

	return false
}
