# Node.js Streaming / Multiple Requests

Interview prompt:

> Write a function that can perform streaming for multiple requests or data sources using Node.js/libuv.

Expected concepts:

- Node.js streams process data chunk by chunk instead of loading everything into memory.
- Streams are useful for large files, uploads, downloads, logs, CSV processing, video/audio, and API responses.
- Streams are event-driven and non-blocking.
- libuv/event loop helps Node handle multiple I/O operations concurrently.
- Each request can create its own readable/writable stream.
- `pipe()` connects streams, but `pipeline()` is preferred because it handles errors and cleanup better.
- Backpressure matters: if the consumer is slower than the producer, the producer should slow down instead of growing memory.
- `Promise.all` can run multiple independent stream pipelines concurrently.

Run:

```bash
node node/streams/multiple-streams.js
```

Interview explanation:

> For streaming multiple requests in Node.js, I would use streams so data is processed chunk by chunk. For each request, I can create a readable stream from the source and pipe it to a writable destination using `pipeline`. This avoids loading the full payload into memory. Since Node.js uses non-blocking I/O through libuv, multiple streams can progress concurrently while the event loop schedules callbacks. I would also handle backpressure and errors properly using `stream.pipeline`.

Important:

> Streams are for I/O concurrency and memory efficiency, not CPU parallelism.

If CPU-heavy processing is needed inside the stream, use `worker_threads` or a separate worker service because CPU-heavy work can block the event loop.

Need to practice:

- `fs.createReadStream()`
- `fs.createWriteStream()`
- `stream.pipeline()`
- Transform streams
- Backpressure
- Handling multiple streams with `Promise.all`
