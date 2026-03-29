package leakybucket

import (
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLeakyBucket(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		lb := NewLeakyBucket(5, 1)

		assert.True(t, lb.Allow())
		assert.Equal(t, 1.0, lb.usedCapacity)

		assert.True(t, lb.Allow())
		assert.Equal(t, 2.0, lb.usedCapacity)

		assert.True(t, lb.Allow())
		assert.Equal(t, 3.0, lb.usedCapacity)

		// advance clock by 1 second
		time.Sleep(time.Second)

		// the bucket will leak by 1 and 1 would be added to bucket keeping the usedCapacity at 3s
		assert.True(t, lb.Allow())
		assert.Equal(t, 3.0, lb.usedCapacity)
	})
}

func TestLeakyBucket_Parallel(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		lb := NewLeakyBucket(3, 1)

		wg := sync.WaitGroup{}

		var success, failure int64

		for range 5 {
			wg.Go(func() {
				wg.Add(1)
				defer wg.Done()

				if lb.Allow() {
					atomic.AddInt64(&success, 1)
				} else {
					atomic.AddInt64(&failure, 1)
				}
			})
		}

		wg.Wait()
		assert.Equal(t, 3.0, lb.usedCapacity)

		assert.Equal(t, int64(3), success)
		assert.Equal(t, int64(2), failure)
	})
}

func TestLeakyBucket_Parallel2(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		lb := NewLeakyBucket(3, 1)

		wg := sync.WaitGroup{}

		var success, failure int64

		// Fire 3 requests at t=0 and 2 requests at t=1 and t=2, so due to leak it should allow the staggered requests
		for i := range 5 {
			if i >= 3 {
				time.Sleep(time.Second)
			}

			wg.Go(func() {
				wg.Add(1)
				defer wg.Done()

				if lb.Allow() {
					atomic.AddInt64(&success, 1)
				} else {
					atomic.AddInt64(&failure, 1)
				}
			})
		}

		wg.Wait()

		assert.Equal(t, 3.0, lb.usedCapacity)
		assert.Equal(t, int64(5), success)
	})
}
