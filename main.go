package main

import (
	"fmt"
	"time"
)

func main() {
	bucket := NewTokenBucket(5, 2)

	for i := 0; i < 10; i++ {
		if bucket.Allow() {
			fmt.Println("✅ Request allowed at", time.Now().Format("15:04:05.000"))
		} else {
			fmt.Println("❌ Rate limited at", time.Now().Format("15:04:05.000"))
		}
		time.Sleep(300 * time.Millisecond)
	}
}