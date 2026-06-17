# Node.js worker_threads

Interview prompt:

> How do you handle CPU-heavy tasks in Node.js without blocking the event loop?

Expected concepts:

- async I/O does not need worker threads
- CPU-heavy JavaScript blocks the event loop
- `worker_threads` runs CPU work on separate threads
- child processes are another option

Run when Node.js is installed:

```bash
node node/worker-threads/main.js
```

