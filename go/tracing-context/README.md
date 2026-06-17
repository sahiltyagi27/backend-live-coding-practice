# Distributed Tracing Context Propagation

Interview prompt:

> How would you implement spans across multiple microservices to understand where a request failed?

This is a lightweight simulation of trace propagation.

Expected concepts:

- one request has one trace ID
- each service creates a span
- spans have parent-child relationship
- pass trace context through HTTP/gRPC/message headers
- OpenTelemetry usually handles this in production
- export spans to Jaeger, Tempo, Datadog, New Relic, etc.

Run:

```bash
go run ./go/tracing-context
```

Interview explanation:

> The first service extracts or creates a trace ID. Each service creates a new span with a parent span ID, then propagates the trace context to downstream calls through headers. In production I would use OpenTelemetry instrumentation and export traces to a backend like Jaeger or Tempo.

