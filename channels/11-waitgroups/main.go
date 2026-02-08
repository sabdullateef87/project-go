package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait() // block until all workers call Done
	fmt.Println("all workers complete")
}

// worker simulates a payment-side task (e.g., saving an authorization).
// WaitGroup ensures main waits for all workers before exiting.
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker", id, "starting")
	// pretend work here
	fmt.Println("worker", id, "done")
}
