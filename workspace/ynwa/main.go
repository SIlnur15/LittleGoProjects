package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens         float64
	maxTokens      float64
	refillRate     float64 // токенов в секунду
	lastRefillTime time.Time
	mutex          sync.Mutex
}

func NewTokenBucket(maxTokens, refillRate float64) *TokenBucket {
	return &TokenBucket{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTime).Seconds()
	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.maxTokens {
		tb.tokens = tb.maxTokens
	}
	tb.lastRefillTime = now

	if tb.tokens >= 1 {
		tb.tokens -= 1
		return true
	}
	return false
}

func main() {
	tb := NewTokenBucket(5, 1) // максимум 5 токенов, пополнение 1 токен/сек
	for i := 1; i <= 10; i++ {
		allowed := tb.AllowRequest()
		fmt.Printf("Запрос %d: %v\n", i, allowed)
		time.Sleep(500 * time.Millisecond)
	}
}
