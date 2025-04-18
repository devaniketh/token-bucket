package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

var quotes = []string{
	"Stay hungry, stay foolish.",
	"The best time to plant a tree was 20 years ago. The second best time is now.",
	"Code is like humor. When you have to explain it, it’s bad.",
	"Simplicity is the soul of efficiency.",
}

var buckets = make(map[string]*TokenBucket)
var mu sync.Mutex

func getBucketForIP(ip string) *TokenBucket {
	mu.Lock()
	defer mu.Unlock()

	if bucket, exists := buckets[ip]; exists {
		return bucket
	}

	// Create new bucket: 5 requests max, 1 token every 2 seconds => 5 tokens in 10 seconds
	bucket := NewTokenBucket(5, 1) // 1 token per 2s = 5 per 10s
	buckets[ip] = bucket
	return bucket
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	bucket := getBucketForIP(ip)

	if !bucket.Allow() {
		http.Error(w, "Too Many Requests. Chill out 🧊", http.StatusTooManyRequests)
		return
	}

	quote := quotes[len(r.URL.Path)%len(quotes)]
	fmt.Fprintln(w, quote)
}

func main() {
	http.HandleFunc("/quotes", quoteHandler)

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}