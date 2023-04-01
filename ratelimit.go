package tinyratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate      int
	interval  time.Duration
	numCalls  int
	lastReset time.Time
	mu        sync.Mutex
}

func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:      rate,
		interval:  interval,
		numCalls:  0,
		lastReset: time.Now(),
		mu:        sync.Mutex{},
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	periodElapsed := time.Since(rl.lastReset)

	if periodElapsed >= rl.interval {
		rl.numCalls = 0
		rl.lastReset = time.Now()
	}

	if rl.numCalls >= rl.rate {
		return false
	}

	rl.numCalls++

	return true
}

func (rl *RateLimiter) Interval() time.Duration {
	return rl.interval
}

func (rl *RateLimiter) Reset() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.numCalls = 0
	rl.lastReset = time.Now()
}

func (rl *RateLimiter) NumCalls() int {
	return rl.numCalls
}
