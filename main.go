package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   int
	tokens     int
	fillRate   int
	mutex      sync.Mutex
	lastRefill time.Time
}

func NewTokenBucket(capacity, fillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		fillRate:   fillRate,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) refillTokens() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	tb.lastRefill = now

	tokensToAdd := int(elapsed.Seconds()) * tb.fillRate
	tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.refillTokens()

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

func main() {
	tb := NewTokenBucket(5, 1) // Capacity 5 tokens, 1 token per second

	// Simulate 10 requests
	for i := 0; i < 10; i++ {
		if tb.AllowRequest() {
			fmt.Println("Request allowed!")
		} else {
			fmt.Println("Request denied. No tokens available.")
		}
		time.Sleep(1 * time.Second)
	}
}
