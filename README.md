# Backend Live Coding Practice

Focused drills from the MAI Labs interview debrief.

Goal:

> Convert backend concepts into working code under interview pressure.

---

## Practice Order

1. [Go worker pool with graceful shutdown](go/worker-pool/README.md)
2. [Go rate limiter middleware](go/rate-limiter/README.md)
3. [Distributed tracing context propagation](go/tracing-context/README.md)
4. [Node.js libuv and concurrent async tasks](node/libuv/README.md)
5. [Node.js worker_threads for CPU-heavy work](node/worker-threads/README.md)
6. [Node.js streams and multiple requests](node/streams/README.md)

---

## How To Practice

For each drill:

1. Read the problem statement.
2. Close the solution file.
3. Implement from scratch in a new file.
4. Run it.
5. Compare with the reference.
6. Explain the solution out loud in interview style.

---

## Interview Targets

You should be able to write these from memory:

- fixed worker pool using `jobs` channel and `sync.WaitGroup`
- graceful shutdown using `context`
- rate limiter middleware returning HTTP `429`
- trace ID/span ID propagation through request context
- Node.js concurrent I/O using `Promise.all`
- Node.js CPU work using `worker_threads`
- Node.js streaming using `pipeline` and backpressure
