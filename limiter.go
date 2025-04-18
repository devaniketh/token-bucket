package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity     int
	tokens       int
	fillInterval time.Duration
	lastRefill   time.Time
	mutex        sync.Mutex
}

func NewTokenBucket(capacity int, ratePerSecond int) *TokenBucket {
	return &TokenBucket{
		capacity:     capacity,
		tokens:       capacity,
		fillInterval: time.Second / time.Duration(ratePerSecond),
		lastRefill:   time.Now(),
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	newTokens := int(elapsed / tb.fillInterval)
	if newTokens > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+newTokens)
		tb.lastRefill = now
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.refill()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}