package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Text  string
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan Job)
	results := make(chan Result)

	var wg sync.WaitGroup
	startWorkers(ctx, 3, jobs, results, &wg)

	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			jobs <- Job{ID: i}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("job=%d result=%s\n", result.JobID, result.Text)
	}

	shutdownCtx, stop := context.WithTimeout(context.Background(), 2*time.Second)
	defer stop()
	gracefulShutdown(shutdownCtx, cancel, &wg)
}

func startWorkers(ctx context.Context, count int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	wg.Add(count)
	for workerID := 1; workerID <= count; workerID++ {
		workerID := workerID
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("worker %d: cancelled\n", workerID)
					return
				case job, ok := <-jobs:
					if !ok {
						fmt.Printf("worker %d: jobs closed\n", workerID)
						return
					}
					time.Sleep(100 * time.Millisecond)
					results <- Result{
						JobID: job.ID,
						Text:  fmt.Sprintf("processed by worker %d", workerID),
					}
				}
			}
		}()
	}
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	cancel()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("shutdown complete")
	case <-ctx.Done():
		fmt.Println("shutdown timed out")
	}
}
