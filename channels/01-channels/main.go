package main

import "fmt"

type Payment struct {
	ID     string
	Amount float64
	Status string
}
  
func newPayment(id string, amount float64) *Payment {
	p := Payment{ID: id, Amount: amount, Status: "authorized"}
	return &p
}


func main() {
	ch := make(chan Payment) // unbuffered channel couples producer/consumer pace
	go paymentProducer(ch)

	for m := range ch {
		processPayment(m)
	}
}

// paymentProducer emits a fixed number of payment status events.
// Demonstrates the simplest unbuffered send/receive.
func paymentProducer(out chan<- Payment) {
	defer close(out) // signal no more events
	for i := 1; i <= 5; i++ {
		out <- *newPayment(fmt.Sprintf("payment-%02d", i), float64(i)*100)
	}
}


func validatePayment(p Payment) bool {
	// Placeholder for validation logic (e.g., check amount, card details)
	fmt.Printf("validating payment with id %s\n", p.ID)
	return p.Amount > 0 && p.Status == "authorized"
}

func processPayment(p Payment) {
	validatePayment(p)
	fmt.Printf("Processing payment %s: amount=%.2f\n", p.ID, p.Amount)
}