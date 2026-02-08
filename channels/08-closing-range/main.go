package main

import "fmt"

func main() {
	ch := make(chan string)
	go producer(ch)
	for v := range ch {
		fmt.Println("processed", v)
	}
	fmt.Println("producer closed channel; loop ended")
}

// producer emits a batch of payouts then closes the channel.
func producer(out chan<- string) {
	for i := 0; i < 5; i++ {
		out <- fmt.Sprintf("PAYOUT-%03d", i)
	}
	close(out)
}
