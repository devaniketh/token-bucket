package main

import (
	"testing"
)

func BenchmarkTokenBucket(b *testing.B) {
	tb := NewTokenBucket(100, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tb.Allow()
	}
}