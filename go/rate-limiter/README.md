# Go Rate Limiter Middleware

Interview prompt:

> Implement/design a rate limiter using Go.

This reference uses a simple in-memory fixed-window limiter.

Expected concepts:

- middleware wraps `http.Handler`
- identify client by IP/API key/user ID
- map stores request count per identity
- mutex protects map
- return `429 Too Many Requests`
- production distributed version should use Redis

Run:

```bash
go run ./go/rate-limiter
curl -i http://localhost:8080/
```

Interview explanation:

> This is an in-memory fixed-window limiter. It is fine for one server. In production with multiple instances, I would move counters to Redis and use atomic operations or Lua scripts. For better burst handling, I would use token bucket or sliding window.

