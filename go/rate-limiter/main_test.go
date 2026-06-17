package main

import (
	"testing"
	"time"
)

func TestFixedWindowLimiterAllowsUntilLimit(t *testing.T) {
	limiter := NewFixedWindowLimiter(2, time.Minute)

	if !limiter.Allow("user-1") {
		t.Fatal("first request should be allowed")
	}
	if !limiter.Allow("user-1") {
		t.Fatal("second request should be allowed")
	}
	if limiter.Allow("user-1") {
		t.Fatal("third request should be blocked")
	}
}

func TestFixedWindowLimiterSeparatesClients(t *testing.T) {
	limiter := NewFixedWindowLimiter(1, time.Minute)

	if !limiter.Allow("user-1") {
		t.Fatal("user-1 first request should be allowed")
	}
	if !limiter.Allow("user-2") {
		t.Fatal("user-2 should have a separate limit")
	}
	if limiter.Allow("user-1") {
		t.Fatal("user-1 second request should be blocked")
	}
}
