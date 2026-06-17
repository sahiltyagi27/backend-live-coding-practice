# Node.js libuv Concurrent Tasks

Interview prompt:

> Use Node.js/libuv and perform concurrent tasks.

Expected concepts:

- JavaScript runs on one main thread
- libuv handles async I/O
- libuv thread pool handles some operations like file system, crypto, DNS
- `Promise.all` can run async I/O concurrently
- CPU-heavy work should use `worker_threads` or child processes
- avoid blocking the event loop

Run when Node.js is installed:

```bash
node node/libuv/promise-all.js
node node/libuv/threadpool-crypto.js
```

