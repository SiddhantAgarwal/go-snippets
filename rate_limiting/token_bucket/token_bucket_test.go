package tokenbucket

import (
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenBucket(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		tb := NewTokenBucket(1, 2)

		assert.True(t, tb.Allow())
		assert.True(t, tb.Allow())

		assert.False(t, tb.Allow())
		assert.False(t, tb.Allow())

		time.Sleep(time.Second * 2)
		assert.True(t, tb.Allow())
	})
}

func TestTokenBucketParallel(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		tb := NewTokenBucket(1, 2)

		var (
			allowedCount int64
			wg           sync.WaitGroup
		)

		for range 10 {
			wg.Go(func() {
				if tb.Allow() {
					atomic.AddInt64(&allowedCount, 1)
				}
			})
		}

		wg.Wait()

		assert.Equal(t, int64(2), allowedCount)

		time.Sleep(time.Second * 2)

		if tb.Allow() {
			atomic.AddInt64(&allowedCount, 1)
		}

		assert.Equal(t, int64(3), allowedCount)
	})
}
