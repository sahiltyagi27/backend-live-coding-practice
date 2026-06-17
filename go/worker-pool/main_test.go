package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestWorkerPoolProcessesAllJobs(t *testing.T) {
	ctx := context.Background()
	jobs := make(chan Job)
	results := make(chan Result, 5)

	var wg sync.WaitGroup
	startWorkers(ctx, 2, jobs, results, &wg)

	for i := 1; i <= 5; i++ {
		jobs <- Job{ID: i}
	}
	close(jobs)

	wg.Wait()
	close(results)

	seen := make(map[int]bool)
	for result := range results {
		seen[result.JobID] = true
	}

	if len(seen) != 5 {
		t.Fatalf("expected 5 processed jobs, got %d", len(seen))
	}
}

func TestGracefulShutdownReturns(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan Job)
	results := make(chan Result)

	var wg sync.WaitGroup
	startWorkers(ctx, 1, jobs, results, &wg)

	shutdownCtx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	gracefulShutdown(shutdownCtx, cancel, &wg)
}
