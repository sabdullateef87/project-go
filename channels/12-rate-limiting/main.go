package main

import (
	"fmt"
	"time"
)

func main() {
	tokens := make(chan struct{}, 5) // burst capacity
	for i := 0; i < cap(tokens); i++ {
		tokens <- struct{}{}
	}

	// refill tokens at 5 per second (1 token every 200ms on average)
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default: // bucket full
			}
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case <-tokens:
			sendToGateway(i)
		default:
			fmt.Println("throttled payment", i)
		}
	}
}

// simulate sending a payment authorization to an external gateway
func sendToGateway(id int) {
	fmt.Println("sent auth for payment", id)
}
