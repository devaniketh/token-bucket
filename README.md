# Token Bucket Rate Limiter in Go

A simple implementation of the **Token Bucket Algorithm** for rate limiting in Go. This rate limiter is used to control the rate of requests to a server by limiting the number of requests a user can make within a specified time frame.

## Features
- **Token Bucket Algorithm**: Efficiently manages request rate using tokens.
- **Configurable Capacity and Fill Rate**: Set the number of tokens (capacity) and the rate of token refill.
- **Concurrency Safe**: Uses mutexes to ensure safe access to the token bucket across multiple goroutines.
- **Benchmarking**: High-performance rate limiter with sub-microsecond response times.

## How It Works
The **Token Bucket** algorithm works by:
- **Filling the bucket** with tokens at a fixed rate.
- **Requests are allowed** if there are available tokens in the bucket.
- **If no tokens are available**, requests are denied until more tokens are added.

The bucket has a defined **capacity** (maximum tokens) and a **fill rate** (tokens added per second).

### Example:
- **Capacity**: 5 tokens
- **Fill Rate**: 1 token per second

If a request is made, and there is a token available, it is **allowed**. If no tokens are left, the request is **denied**.

## Setup & Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/token-bucket.git
   cd token-bucket
   ```
2. **Run the application:**

   ```bash
   go run main.go
   ```

3. **Benchmark the implementation:**
   ```bash
   go test -bench .
   ```
**Example**

```bash
package main

import (
	"fmt"
	"time"
)

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
```

## Key Methods:


``NewTokenBucket(capacity, fillRate int):`` Initializes a new TokenBucket with a specified capacity and fill rate.

``AllowRequest() bool:`` Returns true if a request is allowed (token available), and false otherwise.

``refillTokens():`` Refills tokens based on elapsed time.





## Benchmark Results



![Screenshot 2025-04-18 173220](https://github.com/user-attachments/assets/cf31f092-826b-4d17-aa6f-bf446b8b0d42)

