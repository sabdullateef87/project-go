package main

import "fmt"

func main() {
	ch := make(chan string)
	go producer(ch)
	consumer(ch)
}

// producer sends payment IDs; channel is send-only for clarity.
func producer(out chan<- string) {
	for i := 1; i <= 3; i++ {
		out <- fmt.Sprintf("PAY-%03d", i)
	}
	close(out)
}

// consumer reads from a receive-only channel.
func consumer(in <-chan string) {
	for v := range in {
		fmt.Println("consumed", v)
	}
}
