package main

import (
	"fmt"
	"sync"
	"time"
)

// Payment represents a payment submission entering the system.
type Payment struct {
	ID     string
	Amount float64
}

func main() {
	buf := make(chan Payment, 30) // buffer absorbs bursts before workers drain
	go ingest(buf, 80)

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go processor(i, buf, &wg)
	}

	wg.Wait()
	fmt.Println("all payments processed")
}

// ingest produces payments into a buffered channel to absorb bursty arrivals.
// The buffer smooths producer/consumer pace so workers are not blocked on every send.
func ingest(out chan<- Payment, n int) {
	defer close(out)
	for i := 0; i < n; i++ {
		out <- Payment{ID: fmt.Sprintf("PAY-%03d", i), Amount: float64(50 + i)}
		if i%25 == 0 {
			time.Sleep(50 * time.Millisecond) // simulate bursts
		}
	}
}

// processor validates payments from the channel.
func processor(id int, in <-chan Payment, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range in {
		fmt.Printf("processor %d validating %s amount=%.2f\n", id, p.ID, p.Amount)
	}
}
