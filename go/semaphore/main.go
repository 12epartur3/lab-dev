package main

import (
	"context"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	sema := semaphore.NewWeighted(4)
	jobs := make(chan int, 10)
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		close(jobs)
	}()
	ctx := context.Background()
	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		sema.Acquire(ctx, 1)
		go func(j int) {
			defer sema.Release(1)
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			wg.Done()
		}(j)
	}
	wg.Wait()
}
