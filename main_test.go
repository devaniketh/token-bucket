package main

import (
	"testing"
)

func BenchmarkTokenBucket(b *testing.B) {
	tb := NewTokenBucket(5, 1) // Capacity 5 tokens, 1 token per second

	b.ResetTimer() // Reset timer to ignore setup time

	for i := 0; i < b.N; i++ {
		tb.AllowRequest()
	}
}
