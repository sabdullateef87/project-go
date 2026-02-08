package main

import (
	"fmt"
	"time"
)

func main() {
	feedA := make(chan float64)
	feedB := make(chan float64)

	go priceFeed("A", feedA, 150*time.Millisecond)
	go priceFeed("B", feedB, 300*time.Millisecond)

	timeout := time.After(2 * time.Second)

	// Keep selecting until both feeds are closed or timeout
	for feedA != nil || feedB != nil {
		select {
		case p, ok := <-feedA:
			if !ok {
				feedA = nil
				continue
			}
			fmt.Println("A", p)
		case p, ok := <-feedB:
			if !ok {
				feedB = nil
				continue
			}
			fmt.Println("B", p)
		case <-timeout:
			fmt.Println("stopping due to timeout")
			return
		}
	}

	fmt.Println("all feeds closed")
}

// priceFeed sends a few prices then closes its output channel.
func priceFeed(name string, out chan<- float64, delay time.Duration) {
	defer close(out)
	for i := range 5 {
		time.Sleep(delay)
		out <- 100.0 + float64(i)
	}
}
