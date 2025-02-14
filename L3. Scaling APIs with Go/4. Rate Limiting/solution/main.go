package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// TokenBucket implements the Token Bucket rate limiting algorithm
type TokenBucket struct {
	capacity   int           // Maximum capacity of the bucket (max number of tokens)
	tokens     int           // Current number of tokens
	rate       time.Duration // Time to add one token
	lastFilled time.Time     // Last time tokens were added
	mu         sync.Mutex    // To ensure thread safety
}

func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	// Initialize the TokenBucket with a given capacity and rate.
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		rate:       rate,
		lastFilled: time.Now(),
	}
}

// Implement the logic to refill tokens based on the elapsed time and determine whether to allow the request.
func (tb *TokenBucket) Allow() bool {
	// Lock the mutex to prevent other goroutines from accessing shared data.
	tb.mu.Lock()
	// Ensure the mutex is unlocked when this function exits, regardless of how it exits.
	defer tb.mu.Unlock()

	// Calculate how many tokens should be added since the last refill
	now := time.Now()
	elapsed := now.Sub(tb.lastFilled)
	addTokens := int(elapsed / tb.rate)
	tb.tokens += addTokens

	// Ensure the number of tokens doesn't exceed the capacity
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	// Update the last filled time
	tb.lastFilled = now

	// If there are tokens, consume one and allow the request
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	// No tokens available, reject the request
	return false
}

func main() {
	r := gin.Default()

	// Initialize the Token Bucket rate limiter with a burst of 5 requests and 1 request per second rate
	rateLimiter := NewTokenBucket(5, time.Second)

	// Implement rate limiting on the /data route
	r.GET("/data", func(c *gin.Context) {
		if !rateLimiter.Allow() {
			// Too many requests, return 429 status code
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}
		// Handle successful request
		c.JSON(http.StatusOK, gin.H{"message": "Here is your data"})
	})

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
