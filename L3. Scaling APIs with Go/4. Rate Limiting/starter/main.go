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
}

// Allow checks if a request is allowed based on the number of available tokens
// TODO: Implement the logic in this function to refill tokens based on the elapsed time and determine whether to allow the request.
func (tb *TokenBucket) Allow() bool {
	// Lock the mutex to prevent other goroutines from accessing shared data.
	tb.mu.Lock()
	// Ensure the mutex is unlocked when this function exits, regardless of how it exits.
	defer tb.mu.Unlock()

	// TODO: Calculate how many tokens should be added since the last refill

	// TODO: Ensure the number of tokens doesn't exceed the capacity (hint: be sure to also update the last filled time)

	// TODO: If there are tokens, consume one and allow the request. Otherwise, if no tokens are available, reject the request.
}

func main() {
	r := gin.Default()

	// TODO: Initialize the Token Bucket rate limiter with a burst of 5 requests and 1 request per second rate

	// TODO: Implement rate limiting on the /data route.
	// (be sure to handle "too many" requests, as well as success requests) 

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
