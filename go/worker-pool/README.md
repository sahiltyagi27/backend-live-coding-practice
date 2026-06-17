# Go Worker Pool With Graceful Shutdown

Interview prompt:

> Write Go concurrency code to process many tasks using a fixed number of workers and gracefully shut down.

Expected concepts:

- `jobs` channel
- fixed number of worker goroutines
- `sync.WaitGroup`
- `context.Context`
- close jobs when no more work
- workers exit on closed channel or `ctx.Done()`
- wait for in-flight work or timeout

Run:

```bash
go run ./go/worker-pool
```

Interview explanation:

> I create a fixed number of workers reading from a jobs channel. The producer sends jobs and closes the channel. Each worker selects on `ctx.Done()` and `jobs`, so cancellation can stop new work. A `WaitGroup` tracks workers, and shutdown waits with a timeout so the process does not hang forever.

