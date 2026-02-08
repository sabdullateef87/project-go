package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Demonstrates atomic counters for lightweight payment metrics.
func main() {
	var processed uint64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Increment counter without mutex for a fast metric.
			atomic.AddUint64(&processed, 1)
		}()
	}

	wg.Wait()
	fmt.Println("processed payments:", atomic.LoadUint64(&processed))
}
