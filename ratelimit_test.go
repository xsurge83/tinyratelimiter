package tinyratelimiter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xsurge83/tinyratelimiter"
)

func TestRateLimiter_Allow(t *testing.T) {
	t.Parallel()

	ratelimiter := tinyratelimiter.NewRateLimiter(10, time.Second)

	for i := 0; i < 10; i++ {
		assert.True(t, ratelimiter.Allow())
	}
	assert.False(t, ratelimiter.Allow())

	// Wait for the next interval
	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		assert.True(t, ratelimiter.Allow())
	}
	assert.False(t, ratelimiter.Allow())
}

func TestRateLimiter_Reset(t *testing.T) {
	t.Parallel()

	ratelimiter := tinyratelimiter.NewRateLimiter(2, time.Second)

	ratelimiter.Allow()
	ratelimiter.Reset()

	assert.True(t, ratelimiter.Allow())  // first call should be allowed after reset
	assert.True(t, ratelimiter.Allow())  // second call should be allowed
	assert.False(t, ratelimiter.Allow()) // third call should be rejected (exceeds rate limit)
}

func TestRateLimiter_Reset_Interval(t *testing.T) {
	t.Parallel()

	ratelimiter := tinyratelimiter.NewRateLimiter(2, time.Second)

	ratelimiter.Allow()

	time.Sleep(ratelimiter.Interval() / 2)
	assert.True(t, ratelimiter.Allow())

	assert.False(t, ratelimiter.Allow())

	// Reset the limiter and ensure that we can make another call immediately
	ratelimiter.Reset()
	assert.True(t, ratelimiter.Allow())
}

func TestRateLimiter_Reset_MultiThreaded(t *testing.T) {
	var (
		ratelimiter = tinyratelimiter.NewRateLimiter(2, time.Second)
		done        = make(chan bool)
	)

	t.Parallel()

	// Start two goroutines that will hit the rate limiter concurrently
	go func() {
		ratelimiter.Allow()
		done <- true
	}()

	go func() {
		ratelimiter.Allow()
		done <- true
	}()

	// Wait for both goroutines to finish and check the rate limiter state
	<-done
	<-done

	assert.Equal(t, 2, ratelimiter.NumCalls())

	// Reset the limiter and ensure that it's in the expected state
	ratelimiter.Reset()
	assert.Equal(t, 0, ratelimiter.NumCalls())
	assert.True(t, ratelimiter.Allow())
}
