package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	tokens    int
	maxTokens int
	rate      time.Duration
	mutex    sync.Mutex
}

func NewRateLimiter(maxTokens int, ratePerSecond int) *RateLimiter {
	rl := &RateLimiter{
		tokens:    maxTokens,
		maxTokens: maxTokens,
		rate:      time.Second / time.Duration(ratePerSecond),
	}
	go rl.refill()
	return rl
}

func (rl *RateLimiter) refill() {
	ticker := time.NewTicker(rl.rate)
	defer ticker.Stop()
	for {
		<-ticker.C
		rl.mutex.Lock()
		if rl.tokens < rl.maxTokens {
			rl.tokens++
		}
		rl.mutex.Unlock()
	}
}

func (rl *RateLimiter) AllowRequest() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

type intSubRef struct {
	subID        uint64
	componentUID string
}

func (i *intSubRef) GetComponentUID() string {
	return i.componentUID
}

func (i *intSubRef) GetSubscriptionUID() uint64 {
	return i.subID
}

type internalSubscription struct {
	mapComponent map[string]bool
	mapMutex     sync.Mutex
}

func (i *internalSubscription) isEmpty() bool {
	i.mapMutex.Lock()
	defer i.mapMutex.Unlock()
	return len(i.mapComponent) == 0
}

func rateLimitedHandler(rl *RateLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if rl.AllowRequest() {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Request allowed")
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintln(w, "Too many requests. Slow down!")
		}
	}
}

func main() {
	rlimiter := NewRateLimiter(5, 2)
	http.HandleFunc("/api", rateLimitedHandler(rlimiter))
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
