package main

import (
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
