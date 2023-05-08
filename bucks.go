package bucks

import (
	"sync"
	"time"
)

type nowFunction func() time.Time

type TokenBucket struct {
	capacity   int
	tokens     int
	refillRate int
	lastRefill time.Time
	mutex      sync.Mutex
	now        nowFunction
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
		mutex:      sync.Mutex{},
		now:        time.Now,
	}
}

func (tb *TokenBucket) refill() {
	now := tb.now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += int(elapsed * float64(tb.refillRate))
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

func (tb *TokenBucket) TakeToken(numTokens int) bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.refill()

	if tb.tokens > 0 {
		tb.tokens -= numTokens
		return true
	}

	return false
}
