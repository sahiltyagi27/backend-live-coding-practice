package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type traceContextKey struct{}

type TraceContext struct {
	TraceID string
	SpanID  string
}

func main() {
	ctx := context.Background()
	ctx = startRootSpan(ctx, "api-gateway")

	callService(ctx, "user-service")
	callService(ctx, "order-service")
}

func startRootSpan(ctx context.Context, service string) context.Context {
	trace := TraceContext{
		TraceID: newID("trace"),
		SpanID:  newID("span"),
	}
	fmt.Printf("%s trace=%s span=%s parent=<none>\n", service, trace.TraceID, trace.SpanID)
	return context.WithValue(ctx, traceContextKey{}, trace)
}

func callService(ctx context.Context, service string) {
	parent := ctx.Value(traceContextKey{}).(TraceContext)
	child := TraceContext{
		TraceID: parent.TraceID,
		SpanID:  newID("span"),
	}

	fmt.Printf("%s trace=%s span=%s parent=%s\n", service, child.TraceID, child.SpanID, parent.SpanID)

	// In real HTTP/gRPC code, child.TraceID and child.SpanID would be injected
	// into outbound request headers so the next service can continue the trace.
	_ = context.WithValue(ctx, traceContextKey{}, child)
}

func newID(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano()+rand.Int63n(1000))
}
