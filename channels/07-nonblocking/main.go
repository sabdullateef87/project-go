package main

import "fmt"

// Demonstrates non-blocking send/receive for best-effort payment metrics logging.
func main() {
	metrics := make(chan string, 1)
	metrics <- "payment_approved"

	// non-blocking send: don't block the payment path if metrics queue is full
	select {
	case metrics <- "payment_declined":
		fmt.Println("logged decline")
	default:
		fmt.Println("metrics channel full; skip log")
	}

	// non-blocking receive: flush if available
	select {
	case v := <-metrics:
		fmt.Println("flushed metric", v)
	default:
		fmt.Println("no metric to flush")
	}
}
