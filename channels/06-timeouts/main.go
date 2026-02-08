package main

import (
	"fmt"
	"time"
)

func main() {
	resp := make(chan string, 1)
	go gatewayAuth(1500*time.Millisecond, resp) // slow gateway

	select {
	case status := <-resp:
		fmt.Println("gateway status", status)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout; fall back to async review queue")
	}
}

// gatewayAuth simulates a payment gateway authorization that may be slow.
func gatewayAuth(delay time.Duration, out chan<- string) {
	time.Sleep(delay)
	out <- "approved"
}
