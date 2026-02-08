package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Payment represents an incoming payment authorization request.
type Payment struct {
	ID        string
	Amount    float64
	Currency  string
	CardLast4 string
}

func main() {
	jobs := make(chan Payment)
	results := make(chan string)

	var wg sync.WaitGroup
	numWorkers := 4

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// producer: simulate burst of payment authorizations
	go func() {
		for i := 0; i < 12; i++ {
			amt := 50 + float64(i*120)
			jobs <- Payment{ID: fmt.Sprintf("PAY-%03d", i), Amount: amt, Currency: "USD", CardLast4: fmt.Sprintf("%04d", 1000+i)}
		}
		close(jobs)
	}()

	// collector: close results when workers finish
	go func() { wg.Wait(); close(results) }()

	for r := range results {
		fmt.Println(r)
	}
}

// worker performs basic validation (amount limit + random fraud flag) and returns a decision string.
func worker(id int, jobs <-chan Payment, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	randSrc := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))
	for p := range jobs {
		if p.Amount <= 0 {
			results <- fmt.Sprintf("worker %d DECLINE %s invalid amount", id, p.ID)
			continue
		}
		if p.Amount > 5000 {
			results <- fmt.Sprintf("worker %d REVIEW %s high amount %.2f", id, p.ID, p.Amount)
			continue
		}
		if randSrc.Float64() < 0.05 {
			results <- fmt.Sprintf("worker %d DECLINE %s suspected fraud", id, p.ID)
			continue
		}
		results <- fmt.Sprintf("worker %d APPROVE %s %.2f %s", id, p.ID, p.Amount, p.Currency)
	}
}
