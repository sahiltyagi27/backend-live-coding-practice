package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func main() {
	limiter := NewFixedWindowLimiter(5, time.Minute)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	handler := limiter.Middleware(mux)

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

type clientWindow struct {
	count      int
	windowEnds time.Time
}

type FixedWindowLimiter struct {
	limit  int
	window time.Duration
	mu     sync.Mutex
	hits   map[string]clientWindow
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:  limit,
		window: window,
		hits:   make(map[string]clientWindow),
	}
}

func (l *FixedWindowLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := clientKey(r)
		if !l.Allow(key) {
			w.Header().Set("Retry-After", "60")
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (l *FixedWindowLimiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	window := l.hits[key]

	if now.After(window.windowEnds) {
		window = clientWindow{
			count:      0,
			windowEnds: now.Add(l.window),
		}
	}

	if window.count >= l.limit {
		l.hits[key] = window
		return false
	}

	window.count++
	l.hits[key] = window
	return true
}

func clientKey(r *http.Request) string {
	if apiKey := r.Header.Get("X-API-Key"); apiKey != "" {
		return apiKey
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
