package bucks

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var countTime int

func mockNow() time.Time {
	duration := time.Duration(countTime) * time.Second
	countTime += 1
	return time.Unix(0, duration.Nanoseconds())
}

func TestTakeToken(t *testing.T) {
	tb := NewTokenBucket(3, 1)
	tb.now = mockNow
	tb.lastRefill = tb.now()

	assert.Equal(t, 3, tb.tokens, "Number of tokens should be 3")

	// Starts with 3, ends with 2 -> true
	assert.True(t, tb.TakeToken(1), "Bucket with 3 tokens should allow request 1 token")
	// Starts with 2, refill +1, ends with 1 -> true
	assert.Equal(t, 2, tb.tokens, "Number of tokens should be 2")
	assert.True(t, tb.TakeToken(2), "Bucket with 3 tokens should allow request 2 tokens ")
	// Starts with 1, refill +1, ends with 0 -> true
	assert.Equal(t, 1, tb.tokens, "Number of tokens should be 1")
	assert.True(t, tb.TakeToken(2), "Bucket with 2 tokens should allow request 2 tokens ")
	// Starts with 0, refill +1, ends with 1 -> true
	assert.Equal(t, 0, tb.tokens, "Number of tokens should be 0")
	assert.True(t, tb.TakeToken(2), "Bucket with 1 tokens should allow last request 2 tokens ")
	// Starts with -1, refill +1, ends with 0 -> false
	assert.Equal(t, -1, tb.tokens, "Number of tokens should be 0")
	assert.False(t, tb.TakeToken(2), "Bucket with 0 tokens should NOT allow request 2 tokens ")
	// Starts with 1, refill +1, ends with 0 -> true
	assert.Equal(t, 0, tb.tokens, "Number of tokens should be 1")
	assert.True(t, tb.TakeToken(2), "Bucket with 1 token should allow request 2 tokens")
	assert.Equal(t, -1, tb.tokens, "Number of tokens should be -1")
}
